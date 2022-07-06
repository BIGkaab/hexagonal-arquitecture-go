package mapper

import (
	"github.com/BIGKaab/hexagonal-arquitecture-go/domain"
	"github.com/BIGKaab/hexagonal-arquitecture-go/infraestructure/inside/dto"
	"github.com/BIGKaab/hexagonal-arquitecture-go/infraestructure/outside/gorm/entity"
)

type TaskMapper interface {
	TaskDtoToDomain(task dto.Task) domain.Task
	TaskDomainToEntity(task domain.Task) entity.Task
	TaskEntityToDomain(task entity.Task) domain.Task
	TaskDomainToDto(task domain.Task) dto.Task
}
