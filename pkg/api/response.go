package api

type SimpleResponse struct {
	Message string `json:"message,omitempty"`
	Error   error  `json:"error,omitempty"`
}
