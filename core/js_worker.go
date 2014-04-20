package core

/*
import (
	"github.com/robertkrimen/otto"
	"io/ioutil"
)

func jsWorker(a *AddonManager) func(otto.FunctionCall) otto.Value {
	return func(call otto.FunctionCall) otto.Value {
		args := call.ArgumentList
		if len(args) < 2 {
			a.Log.Error("Not enough arguments to call go")
			return otto.NullValue()
		}

		filepaths := args[0].Object()
		callback := args[1]
		go func() {
			vm := otto.New()

			console, err := vm.Get("console")
			if err != nil {
				a.Log.Error("Error getting console")
			}

			cobj := console.Object()
			cobj.Set("log", jsLog(a))
			cobj.Set("debug", jsDebug(a))
			cobj.Set("error", jsErr(a))

			keys := filepaths.Keys()

			var val otto.Value // The returning value
			for i := 0; i < len(keys); i++ {
				key := keys[i]

				path, err := filepaths.Get(key)
				if err != nil {
					a.Log.Error("Error running worker: ", err)
					break
				}

				file, err := ioutil.ReadFile(path.String())
				if err != nil {
					a.Log.Error("Error running worker: ", path.String(), err)
					continue
				}
				val, err = vm.Run(file)
				if err != nil {
					a.Log.Error("Error running worker: ", path.String(), err)
					continue
				}
			}

			evt := CustomEvent{
				callback,
				[]otto.Value{val},
			}
			a.ScriptCutomEventChan <- evt
		}()
		return otto.TrueValue()
	}
}

// Broken with current otto
func jsImport(a *AddonManager, vm *otto.Otto) func(otto.FunctionCall) otto.Value {
	return func(call otto.FunctionCall) otto.Value {
		args := call.ArgumentList

		for i := 0; i < len(args); i++ {
			arg := args[i]
			filename := arg.String()
			file, err := ioutil.ReadFile(filename)
			if err != nil {
				a.Log.Error("Error importing script", filename, err)
				continue
			}
			_, err = call.Otto.Run(file)
			if err != nil {
				a.Log.Error("Error importing script", filename, err)
				continue
			}
		}
		return otto.TrueValue()
	}
}
*/
