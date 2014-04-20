package core

import (
	"code.google.com/p/goprotobuf/proto"
	messages "github.com/jonas747/fortia-messages"
	"github.com/jonas747/fortia-server/netengine"
)

type netHandler interface{}

// Handles Fortia.sendMessage() messages(client->server)
func netClientServerMsgHandler(e *Engine) netHandler {
	return func(msg messages.FortiaMessage, conn netengine.Connection) {
		playerId, exists := conn.GetSessionData().Get("id")
		if !exists {
			e.Log.Warn("Player id does not exist. client tried to send client->server message")
			return
		}
		playerIdInt, ok := playerId.(int)
		if !ok {
			e.Log.Error("Failed casting interface{} to int, id")
			return
		}

		player := e.Players[playerIdInt]

		evt := NewGeneralEvent(msg.GetName(), msg.GetData(), player.Map())
		e.EmitEvent(evt)
	}
}

func netHelloHandler(e *Engine) netHandler {
	return func(msg messages.Hello, conn netengine.Connection) []byte {
		// Assign id, create e player instance and add it to addons.Players, send response
		id := <-e.idChan
		conn.GetSessionData().Set("id", id)

		player := Player{
			0, 0, 0,
			conn,
			id,
		}
		e.Players[id] = &player

		// Send player connect event
		evt := NewGeneralEvent("playerconnect", player.Map())
		e.EmitEvent(evt)

		response := new(messages.HelloResp)
		response.Message = proto.String("OK")
		response.PlayerId = proto.Int32(int32(id))
		encoded, _ := netengine.EncodeMessage(response, int32(messages.MessageTypes_HELLO_RESP))
		return encoded
	}
}

// TODO: only download new resources if updated versions
func netGetClResources(e *Engine) netHandler {
	return func(msg messages.GetCLResources, conn netengine.Connection) []byte {
		appendResourceScript := func(dest []*messages.Resource, script *Script) []*messages.Resource {
			e.Log.Debug("Sending script ", script.Path)
			res := new(messages.Resource)
			res.Name = proto.String(script.Path)
			test := messages.ResourceType_RType_Script
			res.Type = &test
			res.Script = proto.String(string(script.Source))
			dest = append(dest, res)
			return dest
		}

		// get all resources
		resources := make([]*messages.Resource, 0)
		for _, addon := range e.Addons {
			for _, script := range addon.SharedScripts {
				resources = appendResourceScript(resources, script)
			}
			for _, script := range addon.ClientScripts {
				resources = appendResourceScript(resources, script)
			}
		}
		response := new(messages.GetCLResourcesResp)
		response.Resources = resources
		encoded, err := netengine.EncodeMessage(response, int32(messages.MessageTypes_GETCLRESOURCESRESP))
		if err != nil {
			e.Log.Error("Error encoding GetClResources response ", err)
			return make([]byte, 0)

		}
		return encoded
	}
}

func netClientReady(e *Engine) netHandler {
	return func(msg messages.Empty, conn netengine.Connection) {
		// Emit the event playerjoin
		playerId, exists := conn.GetSessionData().Get("id")
		if !exists {
			e.Log.Warn("Player id does not exist. client tried to send clientReady message")
			return
		}
		playerIdInt, ok := playerId.(int)
		if !ok {
			e.Log.Error("Failed casting interface{} to int, player id")
			return
		}

		player := e.Players[playerIdInt]
		evt := NewGeneralEvent("playerjoin", player.Map())
		e.EmitEvent(evt)

		e.Log.Debug("A player connection is now ready!")
	}
}

func netPlayerMove(e *Engine) netHandler {
	return func(msg messages.PlayerMove, conn netengine.Connection) {
		playerId, exists := conn.GetSessionData().Get("id")
		if !exists {
			e.Log.Warn("Player id does not exist. client tried to send PlayerMove message")
			return
		}
		playerIdInt, ok := playerId.(int)
		if !ok {
			e.Log.Error("Failed casting interface{} to int, id")
			return
		}

		player := e.Players[playerIdInt]
		player.x = float64(*msg.NewX)
		player.y = float64(*msg.NewY)
		player.z = float64(*msg.NewZ)
		e.Players[playerIdInt] = player

		evt := NewGeneralEvent("playermove", player.Map())
		e.EmitEvent(evt)
	}
}
