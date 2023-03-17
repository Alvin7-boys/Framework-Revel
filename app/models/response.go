package models

type Success struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSuccess(message string, data interface{}) Success {
	return Success{
		Message: message,
		Data:    data,
	}
}
