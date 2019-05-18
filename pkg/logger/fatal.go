package logger

import (
	"fmt"
	"os"
)

func Fatal(message string, err error) {

	print(TagFatal, fmt.Sprintf(message, err))
	os.Exit(1)

}
