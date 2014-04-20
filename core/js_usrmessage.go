package core

import (
	"code.google.com/p/goprotobuf/proto"
	"github.com/idada/v8.go"
	messages "github.com/jonas747/fortia-messages"
	"github.com/jonas747/fortia-server/netengine"
)

// Sends a message to a player
func jsUsrMessage(e *Engine) interface{} {
	return func(name string, data *v8.Value, player *v8.Object) {

		if player == nil || data == nil || name == "" {
			log.Error("Tried calling _sendUsrMessage with invalid arguments")
			return
		}
		// Get the player id
		id := player.GetProperty("id").ToInteger()

		RealPlayer, ok := e.Players[int(id)]
		if !ok {
			e.Log.Error("No player found for player id ", id)
			e.Log.Error(string(v8.ToJSON(player.Value)))
			return
		}
		if !RealPlayer.conn.Open() {
			e.Log.Error("Tried to send usermessage to e disconnected player")
			return
		}

		serializedData := string(v8.ToJSON(data))

		msg := new(messages.FortiaMessage)
		msg.Name = proto.String(name)
		msg.Data = proto.String(serializedData)
		serializedMessage, err := netengine.EncodeMessage(msg, int32(messages.MessageTypes_SERVERCLIENTMESSAGE))
		if err != nil {
			e.Log.Error("Error serializing message")
			return
		}
		err = RealPlayer.conn.Send(serializedMessage)
		if err != nil {
			e.Log.Error("Error sending usermessage: ", err)
			return
		}
		return
	}
}
