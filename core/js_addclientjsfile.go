package core

import (
//"github.com/idada/v8.go"
)

func jsAddClientJsFile(e *Engine) interface{} {
	return func(filepath string, exec bool) {
		script, err := LoadScript(e.currentlyLoading.Path+"/scripts/"+filepath, SCRIPT_CLIENT)
		if err != nil {
			e.Log.Error("Error adding client js file [", e.currentlyLoading.Path+"/"+filepath+"]: ", err)
			return
		}
		script.ShouldExec = exec
		e.currentlyLoading.ClientScripts = append(e.currentlyLoading.ClientScripts, script)
	}
}
