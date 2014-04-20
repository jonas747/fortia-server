package core

import (
//v8 "github.com/idada/v8.go"
)

// Returns all the connected players
// TODO: avoid race conditions
func jsGetPlayers(e *Engine) interface{} {
	return func() (ids []int, x []float64, y []float64, z []float64) {
		players := e.Players

		for _, ply := range players {
			ids = append(ids, ply.id)
			x = append(x, ply.x)
			y = append(y, ply.y)
			z = append(z, ply.z)
		}
		return
	}
}
