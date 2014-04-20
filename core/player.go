package core

import (
	"github.com/idada/v8.go"
	"github.com/jonas747/fortia-server/netengine"
)

type Player struct {
	conn netengine.Connection
	id   int
}

func (e *Engine) GetJsPlayer(id int) *v8.Value {
	var retVal *v8.Value
	e.JsContext.Scope(func(cs v8.ContextScope) {
		FortiaObj := cs.Global().GetProperty("Fortia")
		PlayersObj := FortiaObj.ToObject().GetProperty("playerList")
		player := PlayersObj.ToArray().GetElement(id)
		retVal = player
	})
	return retVal
}
