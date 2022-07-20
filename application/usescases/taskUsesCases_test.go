package usescases

import (
	"github.com/BIGKaab/hexagonal-arquitecture-go/dummy"
	"github.com/BIGKaab/hexagonal-arquitecture-go/mocks"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInGetAllTasks_OK(t *testing.T) {
	//mocks
	repo := &mocks.TaskPortOut{}
	repo.On("OutGetAllTasks").Return(dummy.TasksEntity, nil)
	mapper := &mocks.TaskMapper{}
	mapper.On("TaskEntityToDomain", dummy.TaskEntity).Return(dummy.TaskDomain)
	service := NewTaskPortOut(repo, mapper)
	//assert
	tasks, err := service.InGetAllTasks()
	if err != nil {
		log.Error(err)
	}
	assert.NotNil(t, tasks)
	assert.Equal(t, dummy.TasksDomain, tasks)
}

func TestInAddTask_OK(t *testing.T) {
	//mocks
	repo := &mocks.TaskPortOut{}
	repo.On("OutAddTask", &dummy.TaskEntity).Return(&dummy.TaskEntity, nil)
	mapper := &mocks.TaskMapper{}
	mapper.On("TaskDomainToEntity", dummy.TaskDomain).Return(dummy.TaskEntity)
	mapper.On("TaskEntityToDomain", dummy.TaskEntity).Return(dummy.TaskDomain)
	service := NewTaskPortOut(repo, mapper)
	//assert
	task, err := service.InAddTask(dummy.TaskDomain)
	if err != nil {
		log.Error(err)
	}
	assert.NotNil(t, task)
	assert.Equal(t, dummy.TaskDomain, task)
}

func TestInFindTaskById_OK(t *testing.T) {
	//mocks
	repo := &mocks.TaskPortOut{}
	repo.On("OutFindTaskById", 1).Return(dummy.TaskEntity, nil)
	mapper := &mocks.TaskMapper{}
	mapper.On("TaskEntityToDomain", dummy.TaskEntity).Return(dummy.TaskDomain)
	service := NewTaskPortOut(repo, mapper)
	//assert
	task, err := service.InFindTaskById(1)
	if err != nil {
		log.Error(err)
	}
	assert.NotNil(t, task)
	assert.Equal(t, dummy.TaskDomain, task)
}

func TestInUpdateTask_OK(t *testing.T) {
	//mocks
	repo := &mocks.TaskPortOut{}
	repo.On("OutUpdateTask", 1, &dummy.TaskEntity).Return(&dummy.TaskEntity, nil)
	mapper := &mocks.TaskMapper{}
	mapper.On("TaskDomainToEntity", dummy.TaskDomain).Return(dummy.TaskEntity)
	mapper.On("TaskEntityToDomain", dummy.TaskEntity).Return(dummy.TaskDomain)
	service := NewTaskPortOut(repo, mapper)
	//assert
	task, err := service.InUpdateTask(1, dummy.TaskDomain)
	if err != nil {
		log.Error(err)
	}
	assert.NotNil(t, task)
	assert.Equal(t, dummy.TaskDomain, task)
}

func TestDeleteTask_OK(t *testing.T) {
	//mocks
	repo := &mocks.TaskPortOut{}
	repo.On("OutDeleteTask", 1).Return(nil)
	mapper := &mocks.TaskMapper{}
	service := NewTaskPortOut(repo, mapper)
	//assert
	err := service.InDeleteTask(1)
	if err != nil {
		log.Error(err)
	}
	assert.Nil(t, err)
}
