package core

import (
	"github.com/jonas747/fortia-server/netengine"
)

type Player struct {
	x, y, z float64
	conn    netengine.Connection
	id      int
}

// Returns the otto.value player
func (p *Player) Map() map[string]interface{} {
	vals := make(map[string]interface{})
	vals["x"] = p.x
	vals["y"] = p.y
	vals["z"] = p.z
	vals["id"] = p.id
	return vals
}
