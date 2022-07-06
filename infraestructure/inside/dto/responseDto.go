package dto

type Message struct {
	Message string
	Data    interface{}
}

type MessageError struct {
	Message string
}
