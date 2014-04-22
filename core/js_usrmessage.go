package core

import (
	"bytes"
	"code.google.com/p/goprotobuf/proto"
	"compress/zlib"
	"github.com/idada/v8.go"
	messages "github.com/jonas747/fortia-messages"
	"github.com/jonas747/fortia-server/netengine"
)

// Sends a message to a player
func jsUsrMessage(e *Engine) interface{} {
	return func(player *v8.Value, name string, data *v8.Value, compress bool) {

		if player == nil || data == nil || name == "" {
			log.Error("Tried calling _sendUsrMessage with invalid arguments")
			return
		}
		// Get the player id
		id := 0
		if player.IsInt32() || player.IsUint32() {
			id = int(player.ToInteger())
		} else if player.IsObject() {
			id = int(player.ToObject().GetProperty("id").ToInteger())
		} else {
			e.Log.Error("Unsuported player argument type calling jsUsrMessage: ", player.String())
			return
		}

		RealPlayer, ok := e.Players[int(id)]
		if !ok {
			e.Log.Error("No player found for player id ", id)
			e.Log.Error(string(v8.ToJSON(player)))
			return
		}
		if !RealPlayer.conn.Open() {
			e.Log.Error("Tried to send usermessage to e disconnected player")
			return
		}

		serializedData := v8.ToJSON(data)

		msg := new(messages.FortiaMessage)
		// Compress it
		if compress {
			var buffer bytes.Buffer
			writeCloser := zlib.NewWriter(&buffer)
			_, err := writeCloser.Write(serializedData)
			if err != nil {
				e.Log.Error("Error compressing usermessage data: ", err)
				return
			}
			err = writeCloser.Close()
			if err != nil {
				e.Log.Error("Error compressing usermessage data: ", err)
				return
			}
			serializedData = buffer.Bytes()
			msg.Compress = proto.Bool(true)
		}

		msg.Name = proto.String(name)
		msg.Data = proto.String(string(serializedData))
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
		e.Log.Debug("Sent usrmessage of size: ", len(serializedMessage))
		return
	}
}
