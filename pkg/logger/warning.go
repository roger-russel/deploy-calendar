package logger

// Warning log
func Warning(any interface{}) {

	switchType(any, TagWarning)

}
