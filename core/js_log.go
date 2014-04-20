package core

func jsLog(v ...interface{}) {
	log.Info(v...)
}
func jsError(v ...interface{}) {
	log.Error(v...)
}
func jsDebug(v ...interface{}) {
	log.Debug(v...)
}
