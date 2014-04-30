package core

import (
	"github.com/idada/v8.go"
)

type jsIncludePair struct {
	ctx    *v8.Context
	engine *v8.Engine
}

func jsInclude(e *Engine, relative bool, pair *jsIncludePair) interface{} {
	return func(filepath string) {

		path := filepath
		if relative {
			path = e.currentlyLoading.Path + "/scripts/" + path
		}

		script, err := LoadScript(path, SCRIPT_SERVER)
		if err != nil {
			e.Log.Error("Error including js file ["+path+"]: ", err)
			return
		}
		script.Compile(pair.engine)
		script.Run(pair.ctx)
	}
}
