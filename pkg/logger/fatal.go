package logger

import (
	"fmt"
	"os"
)

func Fatal(message string, err error) {

	write(TagFatal, fmt.Sprintf(message, err))
	os.Exit(1)

}
