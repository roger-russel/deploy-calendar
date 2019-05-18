package logger

import "fmt"

func Error(message string, err error) {

	write(TagError, fmt.Sprintf(message, err))

}
