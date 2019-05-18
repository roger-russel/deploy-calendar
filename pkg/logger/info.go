package logger

// Info log
func Info(any interface{}) {

	switchType(any, TagInfo)

}
