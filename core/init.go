package core

import (
	v8 "github.com/idada/v8.go"
	"github.com/jonas747/fortia-server/messages"
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
	globalTemplate := e.JsEngine.NewObjectTemplate()

	var pair = new(jsIncludePair)
	globalTemplate.Bind("include", jsInclude(e, true, pair))
	globalTemplate.Bind("addClientJsFile", jsAddClientJsFile(e))

	ctx := e.JsEngine.NewContext(globalTemplate)

	pair.ctx = ctx
	pair.engine = e.JsEngine

	ctx.Scope(func(cs v8.ContextScope) {
		global := cs.Global()

		// Fortia namespace
		fortiaTemplate := e.JsEngine.NewObjectTemplate()

		fortiaTemplate.Bind("_getPlayers", jsGetPlayers(e))
		fortiaTemplate.Bind("_sendUsrMessage", jsUsrMessage(e))
		fortiaTemplate.Bind("_postWorkerMessage", jsPostMessage(e))
		fortiaTemplate.Bind("_newWorker", jsWorker(e))

		jsFileApi(fortiaTemplate)

		fortiaObj := e.JsEngine.MakeObject(fortiaTemplate)
		global.SetProperty("Fortia", fortiaObj, v8.PA_None)

		// console
		consoleTemplate := e.JsEngine.NewObjectTemplate()
		jsLogApi(consoleTemplate)
		consoleObj := e.JsEngine.MakeObject(consoleTemplate)

		global.SetProperty("console", consoleObj, v8.PA_None)
	})
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
