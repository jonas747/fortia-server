package core

import (
	"github.com/idada/v8.go"
)

func jsLog(v ...interface{}) {
	log.Info(v...)
}
func jsError(v ...interface{}) {
	log.Error(v...)
}
func jsDebug(v ...interface{}) {
	log.Debug(v...)
}

func jsLogApi(template *v8.ObjectTemplate) {
	template.Bind("log", jsLog)
	template.Bind("error", jsError)
	template.Bind("debug", jsDebug)
}
