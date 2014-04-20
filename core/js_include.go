package core

import (
//"github.com/idada/v8.go"
)

func jsInclude(e *Engine) interface{} {
	return func(filepath string) {
		script, err := LoadScript(e.currentlyLoading.Path+"/scripts/"+filepath, SCRIPT_SERVER)
		if err != nil {
			e.Log.Error("Error including js file: ", e.currentlyLoading.Path+"/"+filepath, err)
			return
		}
		script.Compile(e.JsEngine)
		script.Run(e.JsContext)
	}
}
