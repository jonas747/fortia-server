package addons

import (
	"code.google.com/p/goprotobuf/proto"
	"encoding/json"
	"github.com/jonas747/fortia/gameserver/netengine"
	"github.com/jonas747/fortia/messages"
	"github.com/robertkrimen/otto"
)
func (a *AddonManager) InitBaseApi() {
	

	// TODO: clean this up
	a.OttoInstance.Set("_fortiaSendUsrMessage", func(call otto.FunctionCall) otto.Value {
		args := call.ArgumentList
		if len(args) < 3 {
			a.Log.Error("Not enough arguments to call _fortiaSendUsrMessage")
			return otto.NullValue()
		}
		name := call.Argument(0)

		// Get the player id
		player := call.Argument(2)
		id, err := player.Object().Get("id")
		if err != nil {
			a.Log.Error("Error getting id of player")
			return otto.NullValue()
		}

		idInt, err := id.ToInteger()
		if err != nil {
			a.Log.Error("Error getting id of player")
			return otto.NullValue()
		}

		RealPlayer := a.Players[int(idInt)]

		data := call.Argument(1)
		exportedData, err := data.Export()
		if err != nil {
			a.Log.Error("Error exporting data")
			return otto.NullValue()
		}

		serializedData, err := json.Marshal(&exportedData)
		if err != nil {
			a.Log.Error("Error serializing data")
			return otto.NullValue()
		}

		msg := new(messages.FortiaMessage)
		msg.Name = proto.String(name.String())
		msg.Data = proto.String(string(serializedData))
		serializedMessage, err := proto.Marshal(msg)
		if err != nil {
			a.Log.Error("Error serializing message")
			return otto.NullValue()
		}
		RealPlayer.conn.Send(serializedMessage)
		return otto.TrueValue()
	})

	a.OttoInstance.Set("_fortiaRegisterBlockType", func(call otto.FunctionCall) otto.Value {
		args := call.ArgumentList
		if len(args) < 2 {
			a.Log.Error("Not enough arguments to call _fortiaRegisterBlockType")
			return otto.NullValue()
		}

		name := call.Argument(0)
		nameStr := name.String()

		constructor := call.Argument(1)

		id := len(a.BlockTypes)

		btype := BlockType{uint16(id), nameStr, constructor}
		a.BlockTypes = append(a.BlockTypes, btype)
		return otto.TrueValue()
	})
	go func() {
		for {
			disconnected := <-a.NetEngine.ConnCloseChan
			sessionData := disconnected.GetSessionData()
			id, exists := sessionData.Get("id")
			if !exists {
				a.Log.Error("Connection not authorized closed...")
				continue
			}
			idInt := id.(int\)
			ply := a.Players[idInt]
			evt := Event{
				"playerleave",
				ply.ToOttoVal(a.OttoInstance),
			}
			a.ScriptEventChan <- evt
		}
	}()
	a.NetEngine.EmitConnOnClose = true
}
