package impl

import (
	"clean-arquitecture-go/application/mapper"
	"clean-arquitecture-go/domain"
	"clean-arquitecture-go/infraestructure/inside/dto"
	"clean-arquitecture-go/infraestructure/outside/gorm/entity"
)

type TaskMapperImpl struct {
	mapper mapper.TaskMapper
}

func (t TaskMapperImpl) TaskDtoToDomain(task dto.Task) domain.Task {
	return domain.Task{
		Name:        task.Name,
		Description: task.Description,
		Punctuation: task.Punctuation,
	}
}

func (t TaskMapperImpl) TaskDomainToEntity(task domain.Task) entity.Task {
	return entity.Task{
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
