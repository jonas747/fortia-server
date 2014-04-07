package addons

import (
	"github.com/jonas747/fortia/gameserver/netengine"
	"github.com/robertkrimen/otto"
)

type Player struct {
	x, y, z float64
	conn    netengine.Connection
	id      int
}

// Returns the otto.value player
func (p *Player) ToOttoVal(instance *otto.Otto) otto.Value {
	vals := make(map[string]interface{})
	vals["x"] = p.x
	vals["y"] = p.y
	vals["z"] = p.z
	vals["id"] = p.id
	value, _ := instance.ToValue(vals)
	return value
}
