package logger

import (
	"flag"
	"fmt"
	"os"
)

var verbose = flag.Bool("verbose", os.Getenv("LOGGER_VERBOSE") == "true", "print info level logs to stdout")

var lf *os.File

//Init requires full LogPath
func Init(logFileName string) {
	flag.Parse()

	logFullPath := fmt.Sprintf("%s/%s.log", os.Getenv("LOGGER_PATH"), logFileName)

	lf, err := os.OpenFile(logFullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)

	if err != nil {
		fmt.Println("Failed to open log file:", err)
	}

	defer lf.Close()

	Info("Starting Now")

}
