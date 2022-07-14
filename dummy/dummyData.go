package dummy

import (
	"github.com/BIGKaab/hexagonal-arquitecture-go/domain"
	"github.com/BIGKaab/hexagonal-arquitecture-go/infraestructure/inside/dto"
	"github.com/BIGKaab/hexagonal-arquitecture-go/infraestructure/outside/gorm/entity"
)

var (
	TasksDto = []dto.Task{{
		ID:          1,
		Name:        "dummy",
		Description: "Lorem Ipsum is simply dummy text.",
		Punctuation: 9,
	}}
	TaskDto = dto.Task{
		ID:          1,
		Name:        "dummy",
		Description: "Lorem Ipsum is simply dummy text.",
		Punctuation: 9,
	}
	TasksDomain = []domain.Task{{
		ID:          1,
		Name:        "dummy",
		Description: "Lorem Ipsum is simply dummy text.",
		Punctuation: 9,
	}}
	TaskDomain = domain.Task{
		ID:          1,
		Name:        "dummy",
		Description: "Lorem Ipsum is simply dummy text.",
		Punctuation: 9,
	}
	TasksEntity = []entity.Task{{
		ID:          1,
		Name:        "dummy",
		Description: "Lorem Ipsum is simply dummy text.",
		Punctuation: 9,
	}}
	TaskEntity = entity.Task{
		ID:          1,
		Name:        "dummy",
		Description: "Lorem Ipsum is simply dummy text.",
		Punctuation: 9,
	}
)
