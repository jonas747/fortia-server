package addons

import (
	"code.google.com/p/goprotobuf/proto"
	"github.com/jonas747/fortia/gameserver/netengine"
	"github.com/jonas747/fortia/messages"
)

type netHandler interface{}

// Handles Fortia.sendMessage() messages(client->server)
func netClientServerMsgHandler(a *AddonManager) netHandler {
	return func(msg messages.FortiaMessage, conn netengine.Connection) {
		// Emit event
		playerId, exists := conn.GetSessionData().Get("id")
		if !exists {
			a.Log.Warn("Player id does not exist. client tried to send client->server message")
			return
		}
		playerIdInt, ok := playerId.(int)
		if !ok {
			a.Log.Error("Failed casting interface{} to int, id")
			return
		}

		player := a.Players[playerIdInt]

		evt := NetEvent{
			msg.GetName(),
			msg.GetData(),
			player.ToOttoVal(a.OttoInstance),
		}
		a.ScriptEventChan <- evt
	}
}

func netHelloHandler(a *AddonManager) netHandler {
	return func(msg messages.Hello, conn netengine.Connection) []byte {
		// Assign id, create a player instance and add it to addons.Players, send response
		id := <-a.idChan

		conn.GetSessionData().Set("id", id)

		player := Player{
			0, 0, 0,
			conn,
			id,
		}
		a.Players[id] = &player

		response := new(messages.HelloResp)
		response.Message = proto.String("OK")
		encoded, _ := netengine.EncodeMessage(response, int32(messages.MessageTypes_HELLO_RESP))
		return encoded
	}
}

// TODO: only download new resources if updated versions
func netGetClResources(a *AddonManager) netHandler {
	return func(msg messages.GetCLResources, conn netengine.Connection) []byte {
		appendResourceScript := func(dest []*messages.Resource, script *Script) []*messages.Resource {
			a.Log.Debug("Sending script ", script.Path)
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
		for _, addon := range a.Addons {
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
			a.Log.Error("Error encoding GetClResources response ", err)
			return make([]byte, 0)

		}
		return encoded
	}
}

func netClientReady(a *AddonManager) netHandler {
	return func(msg messages.Empty, conn netengine.Connection) {
		// Emit the event playerjoin
		playerId, exists := conn.GetSessionData().Get("id")
		if !exists {
			a.Log.Warn("Player id does not exist. client tried to send clientReady message")
			return
		}
		playerIdInt, ok := playerId.(int)
		if !ok {
			a.Log.Error("Failed casting interface{} to int, id")
			return
		}

		player := a.Players[playerIdInt]

		evt := GeneralEvent{
			"playerjoin",
			player.ToOttoVal(a.OttoInstance),
		}
		a.ScriptEventChan <- evt
		a.Log.Debug("A player connection is now ready!")
	}
}
