package mapper

import (
	"clean-arquitecture-go/domain"
	"clean-arquitecture-go/infraestructure/inside/dto"
	"clean-arquitecture-go/infraestructure/outside/gorm/entity"
)

type TaskMapper interface {
	TaskDtoToDomain(task dto.Task) domain.Task
	TaskDomainToEntity(task domain.Task) entity.Task
	TaskEntityToDomain(task entity.Task) domain.Task
	TaskDomainToDto(task domain.Task) dto.Task
}
