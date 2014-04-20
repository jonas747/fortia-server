package core

/*
TODO: re implement this later

import (
	"github.com/robertkrimen/otto"
)

// registers a blocktype, not sure if this is needed
func jsRegisterBlockType(a *AddonManager) func(otto.FunctionCall) otto.Value {
	return func(call otto.FunctionCall) otto.Value {
		args := call.ArgumentList
		if len(args) < 2 {
			a.Log.Error("Not enough arguments to call _fortiaRegisterBlockType")
			return otto.NullValue()
		}

		name := call.Argument(0)
		nameStr := name.String()

		constructor := call.Argument(1)

		id := len(a.BlockTypes)

		btype := BlockType{uint16(id), nameStr, constructor}
		a.BlockTypes = append(a.BlockTypes, btype)
		return otto.TrueValue()
	}
}
*/
