package controller

import (
	"encoding/json"
	"github.com/BIGKaab/hexagonal-arquitecture-go/dummy"
	"github.com/BIGKaab/hexagonal-arquitecture-go/infraestructure/inside/dto"
	"github.com/BIGKaab/hexagonal-arquitecture-go/infraestructure/inside/enum"
	"github.com/BIGKaab/hexagonal-arquitecture-go/mocks"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	tasks    []dto.Task
	task     dto.Task
	Id       = 1
	IdString = "1"
)

func TestGetAllTasks_OK(t *testing.T) {
	//Setup
	e := echo.New()
	req := httptest.NewRequest(echo.GET, enum.ROUTER_GROUP_GLOBAL, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(enum.ROUTER_TASK_GROUP)
	//Mocks
	mockTask := new(mocks.TaskPortIn)
	mockTask.On("InGetAllTasks").Return(dummy.TasksDomain, nil)
	mockMapper := new(mocks.TaskMapper)
	mockMapper.On("TaskDomainToDto", dummy.TaskDomain).Return(dummy.TaskDto)
	h := NewTaskPortIn(mockTask, mockMapper)
	//Asserts
	if assert.NoError(t, h.GetAllTasks(c)) {
		err := json.Unmarshal(rec.Body.Bytes(), &tasks)
		if err != nil {
			log.Error(err)
		}
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, dummy.TasksDto, tasks)
	}

}

func TestAddTask_OK(t *testing.T) {
	//Setup
	e := echo.New()
	req := httptest.NewRequest(echo.POST, enum.ROUTER_GROUP_GLOBAL, strings.NewReader(dummy.InputTaskJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(enum.ROUTER_TASK_GROUP)
	//Mocks
	mockMapper := new(mocks.TaskMapper)
	mockMapper.On("TaskDtoToDomain", dummy.TaskDto).Return(dummy.TaskDomain)
	mockMapper.On("TaskDomainToDto", dummy.TaskDomain).Return(dummy.TaskDto)
	mockTask := new(mocks.TaskPortIn)
	mockTask.On("InAddTask", dummy.TaskDomain).Return(dummy.TaskDomain, nil)
	h := NewTaskPortIn(mockTask, mockMapper)
	//Asserts
	if assert.NoError(t, h.AddTask(c)) {
		err := json.Unmarshal(rec.Body.Bytes(), &task)
		if err != nil {
			log.Error(err)
		}
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, dummy.TaskDto, task)
	}

}

func TestFindTaskById_OK(t *testing.T) {
	//Setup
	e := echo.New()
	req := httptest.NewRequest(echo.GET, enum.ROUTER_GROUP_GLOBAL, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(enum.ROUTER_TASK_GROUP + enum.ROUTER_PARAM_ID)
	c.SetParamNames(enum.ROUTER_ID)
	c.SetParamValues(IdString)
	//Mocks
	mockMapper := new(mocks.TaskMapper)
	mockMapper.On("TaskDomainToDto", dummy.TaskDomain).Return(dummy.TaskDto)
	mockTask := new(mocks.TaskPortIn)
	mockTask.On("InFindTaskById", Id).Return(dummy.TaskDomain, nil)
	h := NewTaskPortIn(mockTask, mockMapper)
	//Asserts
	if assert.NoError(t, h.FindTaskById(c)) {
		err := json.Unmarshal(rec.Body.Bytes(), &task)
		if err != nil {
			log.Error(err)
		}
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, dummy.TaskDto, task)
	}
}

func TestUpdateTask_OK(t *testing.T) {
	//Setup
	e := echo.New()
	req := httptest.NewRequest(echo.PUT, enum.ROUTER_GROUP_GLOBAL, strings.NewReader(dummy.InputTaskJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(enum.ROUTER_TASK_GROUP + enum.ROUTER_PARAM_ID)
	c.SetParamNames(enum.ROUTER_ID)
	c.SetParamValues(IdString)
	//Mocks
	mockMapper := new(mocks.TaskMapper)
	mockMapper.On("TaskDtoToDomain", dummy.TaskDto).Return(dummy.TaskDomain)
	mockMapper.On("TaskDomainToDto", dummy.TaskDomain).Return(dummy.TaskDto)
	mockTask := new(mocks.TaskPortIn)
	mockTask.On("InUpdateTask", Id, dummy.TaskDomain).Return(dummy.TaskDomain, nil)
	h := NewTaskPortIn(mockTask, mockMapper)
	//Asserts
	if assert.NoError(t, h.UpdateTask(c)) {
		err := json.Unmarshal(rec.Body.Bytes(), &task)
		if err != nil {
			log.Error(err)
		}
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, dummy.TaskDto, task)
	}

}

func TestDeleteTask_OK(t *testing.T) {
	//Setup
	e := echo.New()
	req := httptest.NewRequest(echo.DELETE, enum.ROUTER_GROUP_GLOBAL, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(enum.ROUTER_TASK_GROUP + enum.ROUTER_PARAM_ID)
	c.SetParamNames(enum.ROUTER_ID)
	c.SetParamValues(IdString)
	//Mocks
	mockMapper := new(mocks.TaskMapper)
	//mockMapper.On("TaskDomainToDto", dummy.TaskDomain).Return(dummy.TaskDto)
	mockTask := new(mocks.TaskPortIn)
	mockTask.On("InDeleteTask", Id).Return(nil)
	h := NewTaskPortIn(mockTask, mockMapper)
	//Asserts
	if assert.NoError(t, h.DeleteTask(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}
