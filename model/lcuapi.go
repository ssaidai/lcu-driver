package model

type WatcherMsg struct {
	EventType string `json:"eventType"`
	URI       string `json:"uri"`
	Data      []byte `json:"data"`
}

type LCUErrorResponse struct {
	ErrorCode  string `json:"errorCode"`
	HttpStatus int    `json:"httpStatus"`
	Message    string `json:"message"`
}
