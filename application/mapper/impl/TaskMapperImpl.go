package impl

import (
	"github.com/BIGKaab/hexagonal-arquitecture-go/application/mapper"
	"github.com/BIGKaab/hexagonal-arquitecture-go/domain"
	"github.com/BIGKaab/hexagonal-arquitecture-go/infraestructure/inside/dto"
	"github.com/BIGKaab/hexagonal-arquitecture-go/infraestructure/outside/gorm/entity"
)

type TaskMapperImpl struct {
	mapper mapper.TaskMapper
}

func (t TaskMapperImpl) TaskDtoToDomain(task dto.Task) domain.Task {
	return domain.Task{
		ID:          task.ID,
		Name:        task.Name,
		Description: task.Description,
		Punctuation: task.Punctuation,
	}
}

func (t TaskMapperImpl) TaskDomainToEntity(task domain.Task) entity.Task {
	return entity.Task{
		ID:          task.ID,
		Name:        task.Name,
		Description: task.Description,
		Punctuation: task.Punctuation,
	}
}

func (t TaskMapperImpl) TaskEntityToDomain(task entity.Task) domain.Task {
	return domain.Task{
		ID:          task.ID,
		Name:        task.Name,
		Description: task.Description,
		Punctuation: task.Punctuation,
	}
}

func (t TaskMapperImpl) TaskDomainToDto(task domain.Task) dto.Task {
	return dto.Task{
		ID:          task.ID,
		Name:        task.Name,
		Description: task.Description,
		Punctuation: task.Punctuation,
	}
}

func NewTaskMapperImpl() *TaskMapperImpl {
	return &TaskMapperImpl{mapper: TaskMapperImpl{}}
}
