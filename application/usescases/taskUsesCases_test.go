package usescases

import (
	"github.com/BIGKaab/hexagonal-arquitecture-go/dummy"
	"github.com/BIGKaab/hexagonal-arquitecture-go/infraestructure/inside/enum"
	"github.com/BIGKaab/hexagonal-arquitecture-go/infraestructure/outside/gorm/entity"
	"github.com/BIGKaab/hexagonal-arquitecture-go/mocks"
	"testing"
)

func TestInGetAllTasks(t *testing.T) {
	testCases := []struct {
		Name          string
		ExpectedData  []entity.Task
		ExpectedError error
	}{
		{
			Name:          "OK",
			ExpectedData:  dummy.TasksEntity,
			ExpectedError: nil,
		},
		{
			Name:          "Internal Server Error",
			ExpectedData:  nil,
			ExpectedError: enum.INTERNAL_SERVER_ERROR,
		},
	}

	mapper := &mocks.TaskMapper{}
	mapper.On("TaskEntityToDomain", dummy.TaskEntity).Return(dummy.TaskDomain)

	for i := range testCases {
		tc := testCases[i]
		repo := &mocks.TaskPortOut{}
		repo.On("OutGetAllTasks").Return(tc.ExpectedData, tc.ExpectedError)
		service := NewTaskPortOut(repo, mapper)

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			_, err := service.InGetAllTasks()
			if err != tc.ExpectedError {
				t.Errorf("Expected %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

func TestInAddTask(t *testing.T) {
	testCases := []struct {
		Name          string
		ExpectedData  entity.Task
		ExpectedError error
	}{
		{
			Name:          "OK",
			ExpectedData:  dummy.TaskEntity,
			ExpectedError: nil,
		},
		{
			Name:          "Internal Server Error",
			ExpectedData:  dummy.TaskEntity,
			ExpectedError: enum.INTERNAL_SERVER_ERROR,
		},
	}

	mapper := &mocks.TaskMapper{}
	mapper.On("TaskDomainToEntity", dummy.TaskDomain).Return(dummy.TaskEntity)
	mapper.On("TaskEntityToDomain", dummy.TaskEntity).Return(dummy.TaskDomain)

	for i := range testCases {
		tc := testCases[i]
		repo := &mocks.TaskPortOut{}
		repo.On("OutAddTask", &dummy.TaskEntity).Return(&tc.ExpectedData, tc.ExpectedError)
		service := NewTaskPortOut(repo, mapper)

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			_, err := service.InAddTask(dummy.TaskDomain)
			if err != tc.ExpectedError {
				t.Errorf("Expected %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

func TestInUpdateTask(t *testing.T) {
	testCases := []struct {
		Name          string
		ExpectedData  entity.Task
		ExpectedId    int
		ExpectedError error
	}{
		{
			Name:          "OK",
			ExpectedId:    1,
			ExpectedData:  dummy.TaskEntity,
			ExpectedError: nil,
		},
		{
			Name:          "Internal Server Error",
			ExpectedId:    1,
			ExpectedData:  dummy.TaskEntity,
			ExpectedError: enum.INTERNAL_SERVER_ERROR,
		},
	}

	mapper := &mocks.TaskMapper{}
	mapper.On("TaskDomainToEntity", dummy.TaskDomain).Return(dummy.TaskEntity)
	mapper.On("TaskEntityToDomain", dummy.TaskEntity).Return(dummy.TaskDomain)

	for i := range testCases {
		tc := testCases[i]
		repo := &mocks.TaskPortOut{}
		repo.On("OutUpdateTask", tc.ExpectedId, &dummy.TaskEntity).Return(&tc.ExpectedData, tc.ExpectedError)
		service := NewTaskPortOut(repo, mapper)

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			_, err := service.InUpdateTask(tc.ExpectedId, dummy.TaskDomain)
			if err != tc.ExpectedError {
				t.Errorf("Expected %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

func TestInFindTaskById(t *testing.T) {
	testCases := []struct {
		Name          string
		ExpectedData  entity.Task
		ExpectedId    int
		ExpectedError error
	}{
		{
			Name:          "OK",
			ExpectedId:    1,
			ExpectedData:  dummy.TaskEntity,
			ExpectedError: nil,
		},
		{
			Name:          "Internal Server Error",
			ExpectedId:    1,
			ExpectedData:  dummy.TaskEntity,
			ExpectedError: enum.INTERNAL_SERVER_ERROR,
		},
	}

	mapper := &mocks.TaskMapper{}
	mapper.On("TaskEntityToDomain", dummy.TaskEntity).Return(dummy.TaskDomain)

	for i := range testCases {
		tc := testCases[i]
		repo := &mocks.TaskPortOut{}
		repo.On("OutFindTaskById", tc.ExpectedId).Return(tc.ExpectedData, tc.ExpectedError)
		service := NewTaskPortOut(repo, mapper)

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			_, err := service.InFindTaskById(tc.ExpectedId)
			if err != tc.ExpectedError {
				t.Errorf("Expected %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

func TestInDeleteTaskId(t *testing.T) {
	testCases := []struct {
		Name          string
		ExpectedId    int
		ExpectedError error
	}{
		{
			Name:          "OK",
			ExpectedId:    1,
			ExpectedError: nil,
		},
		{
			Name:          "Internal Server Error",
			ExpectedId:    1,
			ExpectedError: enum.INTERNAL_SERVER_ERROR,
		},
	}

	mapper := &mocks.TaskMapper{}

	for i := range testCases {
		tc := testCases[i]
		repo := &mocks.TaskPortOut{}
		repo.On("OutDeleteTask", tc.ExpectedId).Return(tc.ExpectedError)
		service := NewTaskPortOut(repo, mapper)

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			err := service.InDeleteTask(tc.ExpectedId)
			if err != tc.ExpectedError {
				t.Errorf("Expected %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

//GetAll
/*
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

func TestInGetAllTasks_InternalServerError(t *testing.T) {
	//mocks
	repo := &mocks.TaskPortOut{}
	repo.On("OutGetAllTasks").Return(dummy.TasksEntity, enum.INTERNAL_SERVER_ERROR)
	mapper := &mocks.TaskMapper{}
	service := NewTaskPortOut(repo, mapper)
	//assert
	_, err := service.InGetAllTasks()
	assert.NotNil(t, err)
	assert.Equal(t, enum.INTERNAL_SERVER_ERROR, err)
}

//Add
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

func TestInAddTask_InternalServerError(t *testing.T) {
	//mocks
	repo := &mocks.TaskPortOut{}
	repo.On("OutAddTask", &dummy.TaskEntity).Return(&dummy.TaskEntity, enum.INTERNAL_SERVER_ERROR)
	mapper := &mocks.TaskMapper{}
	mapper.On("TaskDomainToEntity", dummy.TaskDomain).Return(dummy.TaskEntity)
	service := NewTaskPortOut(repo, mapper)
	//assert
	_, err := service.InAddTask(dummy.TaskDomain)
	assert.NotNil(t, err)
	assert.Equal(t, enum.INTERNAL_SERVER_ERROR, err)
}

//FindById

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

func TestInFindTaskById_InternalServerError(t *testing.T) {
	//mocks
	repo := &mocks.TaskPortOut{}
	repo.On("OutFindTaskById", 1).Return(dummy.TaskEntity, enum.INTERNAL_SERVER_ERROR)
	mapper := &mocks.TaskMapper{}
	service := NewTaskPortOut(repo, mapper)
	//assert
	_, err := service.InFindTaskById(1)
	assert.NotNil(t, err)
	assert.Equal(t, enum.INTERNAL_SERVER_ERROR, err)
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

func TestInUpdateTask_InternalServerError(t *testing.T) {
	//mocks
	repo := &mocks.TaskPortOut{}
	repo.On("OutUpdateTask", 1, &dummy.TaskEntity).Return(&dummy.TaskEntity, enum.INTERNAL_SERVER_ERROR)
	mapper := &mocks.TaskMapper{}
	mapper.On("TaskDomainToEntity", dummy.TaskDomain).Return(dummy.TaskEntity)
	service := NewTaskPortOut(repo, mapper)
	//assert
	_, err := service.InUpdateTask(1, dummy.TaskDomain)
	assert.NotNil(t, err)
	assert.Equal(t, enum.INTERNAL_SERVER_ERROR, err)
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

func TestDeleteTask_InternalServerError(t *testing.T) {
	//mocks
	repo := &mocks.TaskPortOut{}
	repo.On("OutDeleteTask", 1).Return(enum.INTERNAL_SERVER_ERROR)
	mapper := &mocks.TaskMapper{}
	service := NewTaskPortOut(repo, mapper)
	//assert
	err := service.InDeleteTask(1)
	assert.NotNil(t, err)
	assert.Equal(t, enum.INTERNAL_SERVER_ERROR, err)
}
*/
