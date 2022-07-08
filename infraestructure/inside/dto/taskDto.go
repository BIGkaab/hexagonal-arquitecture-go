package dto

import "github.com/go-playground/validator/v10"

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name" validate:"required" example:"Estudiar Go"`
	Description string `json:"description" validate:"required" example:"Comprender la arquitectura hexagonal"`
	Punctuation int    `json:"punctuation" validate:"required,min=1,max=10" example:"2"`
}

func (e *Task) Validate() error {
	return validator.New().Struct(e)
}
