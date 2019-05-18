package logger

import "fmt"

// Info log
func Info(any interface{}) {

	switchType(any, TagInfo)

}

func printInfo(log string) {
	fmt.Println(log)
}
