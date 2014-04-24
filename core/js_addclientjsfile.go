package core

import (
//"github.com/idada/v8.go"
)

func jsAddClientJsFile(e *Engine) interface{} {
	return func(filepath string) {
		script, err := LoadScript(e.currentlyLoading.Path+"/scripts/"+filepath, SCRIPT_SERVER)
		if err != nil {
			e.Log.Error("Error adding client js file [", e.currentlyLoading.Path+"/"+filepath+"]: ", err)
			return
		}
		e.currentlyLoading.ClientScripts = append(e.currentlyLoading.ClientScripts, script)
	}
}
