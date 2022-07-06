package usescases

import (
	"clean-arquitecture-go/application/mapper"
	"clean-arquitecture-go/application/mapper/impl"
	"clean-arquitecture-go/application/port/out"
	"clean-arquitecture-go/domain"
	"clean-arquitecture-go/infraestructure/outside/gorm/repo"
	"github.com/labstack/gommon/log"
)

type TaskServiceRepo struct {
	portOut out.TaskPortOut
	mapper  mapper.TaskMapper
}

func (t TaskServiceRepo) InGetAllTasks() ([]domain.Task, error) {
	var tasksDomain []domain.Task

	tasksEntity, err := t.portOut.OutGetAllTasks()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	for _, taskEntity := range tasksEntity {
		tasksDomain = append(tasksDomain, t.mapper.TaskEntityToDomain(taskEntity))
	}

	return tasksDomain, nil
}

func (t TaskServiceRepo) InAddTask(task domain.Task) (domain.Task, error) {
	taskEntity := t.mapper.TaskDomainToEntity(task)
	resDomain, err := t.portOut.OutAddTask(&taskEntity)
	if err != nil {
		log.Error(err)
		return task, err
	}
	task.ID = resDomain.ID
	return task, nil
}

func (t TaskServiceRepo) InFindTaskById(ID int) (domain.Task, error) {
	var task domain.Task
	taskEntity, err := t.portOut.OutFindTaskById(ID)
	if err != nil {
		log.Error(err)
		return task, err
	}
	task = t.mapper.TaskEntityToDomain(taskEntity)
	return task, nil
}

func (t TaskServiceRepo) InUpdateTask(ID int, task domain.Task) (domain.Task, error) {
	taskEntity := t.mapper.TaskDomainToEntity(task)
	resDomain, err := t.portOut.OutInUpdateTask(ID, &taskEntity)
	if err != nil {
		log.Error(err)
		return task, err
	}
	task.ID = resDomain.ID
	return task, nil
}

func (t TaskServiceRepo) InDeleteTask(ID int) error {
	err := t.portOut.OutDeleteTask(ID)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func NewTaskPortOut() *TaskServiceRepo {
	return &TaskServiceRepo{portOut: repo.NewTaskRepo(), mapper: impl.NewTaskMapperImpl()}
}
