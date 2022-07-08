package out

import (
	"github.com/BIGKaab/hexagonal-arquitecture-go/infraestructure/outside/gorm/entity"
)

type TaskPortOut interface {
	OutGetAllTasks() ([]entity.Task, error)
	OutAddTask(task *entity.Task) (*entity.Task, error)
	OutFindTaskById(ID int) (entity.Task, error)
	OutUpdateTask(ID int, task *entity.Task) (*entity.Task, error)
	OutDeleteTask(ID int) error
}
