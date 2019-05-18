package logger

import (
	"encoding/json"
	"fmt"
)

//Level of logs
const (
	TagInfo    = "INFO : "
	TagWarning = "WARN : "
	TagError   = "ERROR: "
	TagFatal   = "FATAL: "
)

type log struct {
	Level   string `json:"level"`
	Message string `json:"message"`
	//	Context interface `json:"context,omitempty"`
}

func print(lvl string, message string) {

	l := log{Level: lvl, Message: message}

	jsonLog, err := json.Marshal(l)

	if err != nil {
		fmt.Println("Impossible convert: ", l)
	}

	fmt.Println(string(jsonLog))

}

func switchType(any interface{}, lvl string) {

	switch v := any.(type) {
	case string:
		print(lvl, fmt.Sprintf("%s", any))

	default:
		fmt.Println(v)
	}

}
