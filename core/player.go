package core

import (
	"github.com/jonas747/fortia-server/netengine"
)

type Player struct {
	conn netengine.Connection
	id   int
}
