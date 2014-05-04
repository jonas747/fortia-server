package core

/*
The handshake:

1. The client connects to the server using websockets
2. The client send a Hello message
3. The server responds with a Hello response
4. When the client is ready(after downloading resources and such) the client send a ClientReady message
*/

import (
	"code.google.com/p/goprotobuf/proto"
	"github.com/idada/v8.go"
	"github.com/jonas747/fortia-server/messages"
	"github.com/jonas747/fortia-server/netengine"
	"reflect"
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

		evt := e.NewNetEvent(msg.GetName(), string(msg.GetData()), playerIdInt)
		e.EmitEvent(evt)
	}
}

func netHelloHandler(e *Engine) netHandler {
	return func(msg messages.Hello, conn netengine.Connection) []byte {
		// Assign id, create e player instance and add it to addons.Players, send response
		id := <-e.plyIdChan
		conn.GetSessionData().Set("id", id)

		player := Player{
			conn,
			id,
		}
		e.Players[id] = &player

		// Create javascript player object
		e.JsContext.Scope(func(cs v8.ContextScope) {
			function := cs.Global().GetProperty("Fortia").ToObject().GetProperty("initPlayer").ToFunction()
			function.Call(e.JsEngine.GoValueToJsValue(reflect.ValueOf(id)))
		})

		// Send player connect event
		evt := NewGeneralEvent("playerconnect", id)
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
			res := new(messages.Resource)
			res.Name = proto.String(script.Path)
			kind := messages.ResourceType_RType_Script
			res.Type = &kind
			res.Script = proto.String(string(script.Source))
			res.ShouldExecSCript = proto.Bool(script.ShouldExec)
			dest = append(dest, res)
			return dest
		}

		// get all resources
		resources := make([]*messages.Resource, 0)
		for _, addon := range e.Addons {
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
		evt := NewGeneralEvent("playerjoin", player.id)
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

		evt := NewGeneralEvent("playermove", playerIdInt, msg.GetNewX(), msg.GetNewY(), msg.GetNewZ())
		e.EmitEvent(evt)
	}
}
