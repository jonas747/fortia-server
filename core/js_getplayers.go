package core

import (
//v8 "github.com/idada/v8.go"
)

// Returns all the connected players
// TODO: avoid race conditions
func jsGetPlayers(e *Engine) interface{} {
	return func() []int {
		players := e.Players

		var ids []int
		for _, ply := range players {
			ids = append(ids, ply.id)
		}
		return ids
	}
}
