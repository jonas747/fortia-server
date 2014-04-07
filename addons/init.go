package addons

import (
	"github.com/jonas747/fortia/gameserver/netengine"
	"github.com/jonas747/fortia/messages"
)

func addHandler(callback interface{}, evt messages.MessageTypes, engine *netengine.Engine) {
	handler, err := netengine.NewHandler(callback, int32(evt))
	if err != nil {
		panic(err)
	}
	engine.AddHandler(handler)
}

func (a *AddonManager) AddHandlers() {
	addHandler(netClientServerMsgHandler(a), messages.MessageTypes_CLIENTSERVERMESSAGE, a.NetEngine)
	addHandler(netHelloHandler(a), messages.MessageTypes_HELLO, a.NetEngine)
	addHandler(netGetClResources(a), messages.MessageTypes_GETCLRESOURCES, a.NetEngine)
}

func (a *AddonManager) AddExtensions() {
	go func() {
		for {
			disconnected := <-a.NetEngine.ConnCloseChan
			sessionData := disconnected.GetSessionData()
			id, exists := sessionData.Get("id")
			if !exists {
				a.Log.Error("Connection not authorized closed...")
				continue
			}
			idInt := id.(int)
			ply := a.Players[idInt]
			evt := Event{
				"playerleave",
				ply.ToOttoVal(a.OttoInstance),
			}
			a.ScriptEventChan <- evt
		}
	}()
	a.OttoInstance.Set("_fortiaSendUsrMessage", jsUsrMessage(a))
	a.OttoInstance.Set("_fortiaRegisterBlockType", jsRegisterBlockType(a))
	a.OttoInstance.Set("_fortialog", jsLog(a))
	a.OttoInstance.Set("_fortiaerror", jsErr(a))
}

func idGen(out chan int) {
	curId := 1
	go func() {
		for {
			out <- curId
			curId++
		}
	}()
}
