package in

import (
	"clean-arquitecture-go/domain"
)

type TaskPortIn interface {
	InGetAllTasks() ([]domain.Task, error)
	InAddTask(task domain.Task) (domain.Task, error)
	InFindTaskById(ID int) (domain.Task, error)
	InUpdateTask(ID int, task domain.Task) (domain.Task, error)
	InDeleteTask(ID int) error
}
