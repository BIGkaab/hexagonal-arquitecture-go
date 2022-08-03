package dto

type Message struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type MessageError struct {
	Message string `json:"message"`
}
