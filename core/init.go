package core

import (
	v8 "github.com/idada/v8.go"
	messages "github.com/jonas747/fortia-messages"
	"github.com/jonas747/fortia-server/netengine"
)

func addHandler(callback interface{}, evt messages.MessageTypes, engine *netengine.Engine) {
	handler, err := netengine.NewHandler(callback, int32(evt))
	if err != nil {
		panic(err)
	}
	engine.AddHandler(handler)
}

func (a *Engine) AddNetHandlers() {
	addHandler(netClientServerMsgHandler(a), messages.MessageTypes_CLIENTSERVERMESSAGE, a.NetEngine)
	addHandler(netHelloHandler(a), messages.MessageTypes_HELLO, a.NetEngine)
	addHandler(netGetClResources(a), messages.MessageTypes_GETCLRESOURCES, a.NetEngine)
	addHandler(netClientReady(a), messages.MessageTypes_CLIENTREADY, a.NetEngine)
	addHandler(netPlayerMove(a), messages.MessageTypes_PLAYERMOVE, a.NetEngine)
}

func (e *Engine) AddJsExtensions() {
	e.JsEngine.AddMessageListener(func(msg *v8.Message) {
		e.Log.Errorf("%s:%d:%d: %s", msg.ScriptResourceName, msg.Line, msg.StartColumn, msg.Message)
	})
	global := e.JsEngine.NewObjectTemplate()

	global.Bind("_logInfo", jsLog)
	global.Bind("_logDebug", jsDebug)
	global.Bind("_logError", jsError)

	global.Bind("_fortiaGetPlayers", jsGetPlayers(e))
	global.Bind("_fortiaSendUsrMessage", jsUsrMessage(e))

	ctx := e.JsEngine.NewContext(global)
	e.JsContext = ctx
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
