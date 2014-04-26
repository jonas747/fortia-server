package core

import (
//"github.com/idada/v8.go"
)

func jsInclude(e *Engine, relative bool) interface{} {
	return func(filepath string) {

		path := filepath
		if relative {
			path = e.currentlyLoading.Path + "/scripts/" + path
		}

		script, err := LoadScript(path, SCRIPT_SERVER)
		if err != nil {
			e.Log.Error("Error including js file [", e.currentlyLoading.Path+"/"+filepath+"]: ", err)
			return
		}
		script.Compile(e.JsEngine)
		script.Run(e.JsContext)
	}
}
