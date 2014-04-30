/*
Simple webworker implmentation, very lackluster

This is a very somplified version of webworkers in that it does not support
nested workers, and it does not use structured cloning (using json for now, not top priority)

*/

package core

import (
	"github.com/idada/v8.go"
)

type worker struct {
	engine  *v8.Engine
	context *v8.Context
}

func jsWorker(e *Engine) interface{} {
	return func(workerPath string) int {
		id := <-e.workerIdChan

		instance := v8.NewEngine()
		instance.AddMessageListener(func(msg *v8.Message) {
			e.Log.Errorf("%s:%d:%d: %s", msg.ScriptResourceName, msg.Line, msg.StartColumn, msg.Message)
		})

		globalTemplate := instance.NewObjectTemplate()

		pair := new(jsIncludePair)

		globalTemplate.Bind("include", jsInclude(e, false, pair))
		ctx := instance.NewContext(globalTemplate)

		pair.ctx = ctx
		pair.engine = instance

		ctx.Scope(func(cs v8.ContextScope) {
			global := cs.Global()

			// Self template
			selfObjTmpl := instance.NewObjectTemplate()
			selfObjTmpl.Bind("postMessage", jsPostMessageFromWorker(e, id))

			selfObj := instance.MakeObject(selfObjTmpl)
			global.SetProperty("self", selfObj, v8.PA_None)

			// console
			consoleTemplate := instance.NewObjectTemplate()
			jsLogApi(consoleTemplate)
			consoleObj := instance.MakeObject(consoleTemplate)

			global.SetProperty("console", consoleObj, v8.PA_None)
		})

		// Load and exec script
		script, err := LoadScript(workerPath, SCRIPT_SERVER)
		if err != nil {
			e.Log.Error("Error spawning worker: ", err)
			return 0
		}
		script.Compile(instance)

		w := worker{
			instance,
			ctx,
		}

		e.workers[id] = w

		go script.Run(ctx)

		return id
	}
}

func jsPostMessageFromWorker(e *Engine, id int) interface{} {
	return func(data *v8.Value) {
		str := string(v8.ToJSON(data))
		go e.JsContext.Scope(func(cs v8.ContextScope) {
			fortiaVal := cs.Global().GetProperty("Fortia")
			fortiaObj := fortiaVal.ToObject()
			msgVal := fortiaObj.GetProperty("_onWorkerMsg")
			msgFunc := msgVal.ToFunction()

			msgFunc.Call(e.ToJsValue(id, cs), cs.ParseJSON(str))
		})
	}
}

func jsPostMessage(e *Engine) interface{} {
	return func(id int, msg *v8.Value) {
		work, found := e.workers[id]
		if !found {
			e.Log.Error("Tried to call postMessage on invalid worker")
			return
		}
		str := string(v8.ToJSON(msg))
		go work.context.Scope(func(cs v8.ContextScope) {
			selfObj := cs.Global().GetProperty("self")
			callback := selfObj.ToObject().GetProperty("onmessage")
			if callback.IsFunction() {

				cfunc := callback.ToFunction()
				cfunc.Call(cs.ParseJSON(str))
			}
		})
	}
}
