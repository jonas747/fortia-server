package addons

import (
	"github.com/robertkrimen/otto"
)

type Block struct {
	x, y, z int
	kind    int
}

type BlockType struct {
	Id          uint16
	Name        string
	Constructor otto.Value
}
