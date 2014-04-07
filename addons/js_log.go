package addons

import (
	"github.com/robertkrimen/otto"
)

func jsLog(a *AddonManager) func(otto.FunctionCall) otto.Value {
	return func(call otto.FunctionCall) otto.Value {
		args := call.ArgumentList
		if len(args) < 1 {
			a.Log.Error("Not enough arguments to call _fortiaRegisterBlockType")
			return otto.NullValue()
		}

		arg1 := args[0]
		exported, err := arg1.Export()
		if err != nil {
			a.Log.Error("Error exporting argument ", err)
			return otto.NullValue()
		}
		a.Log.Info("Script: ", exported)
		return otto.TrueValue()
	}
}

func jsErr(a *AddonManager) func(otto.FunctionCall) otto.Value {
	return func(call otto.FunctionCall) otto.Value {
		args := call.ArgumentList
		if len(args) < 1 {
			a.Log.Error("Not enough arguments to call _fortiaRegisterBlockType")
			return otto.NullValue()
		}

		arg1 := args[0]
		exported, err := arg1.Export()
		if err != nil {
			a.Log.Error("Error exporting argument ", err)
			return otto.NullValue()
		}
		a.Log.Error("Script: ", exported)
		return otto.TrueValue()
	}
}
