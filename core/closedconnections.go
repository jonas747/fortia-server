package core

import (
	"github.com/idada/v8.go"
	"reflect"
)

// Keeps track of all the scripts event and such
func (e *Engine) ListenForClosedConnection() {

	// make sure the netengine emits the closed connections on the connclose channel
	e.NetEngine.EmitConnOnClose = true

	// Goroutine for receiving disconnected player
	// TODO: avoid race conditions, make it possible to stop
	for {
		disconnected := <-e.NetEngine.ConnCloseChan
		sessionData := disconnected.GetSessionData()
		id, exists := sessionData.Get("id")
		if !exists {
			e.Log.Error("Connection not authorized closed...")
			continue
		}

		idInt := id.(int)
		evt := NewGeneralEvent("playerleave", idInt)

		e.EmitEvent(evt)

		delete(e.Players, idInt)

		e.JsContext.Scope(func(cs v8.ContextScope) {
			function := cs.Global().GetProperty("Fortia").ToObject().GetProperty("removePlayer").ToFunction()
			function.Call(e.JsEngine.GoValueToJsValue(reflect.ValueOf(idInt)))
		})
	}
}
