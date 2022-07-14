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
	"testing"
)

func TestGetAllTasks(t *testing.T) {
	var data []dto.Task
	//Setup
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(enum.ROUTER_GROUP_GLOBAL + enum.ROUTER_TASK_GROUP)
	//Mocks
	mockTask := new(mocks.TaskPortIn)
	mockTask.On("InGetAllTasks").Return(dummy.TasksDomain, nil)
	mockMapper := new(mocks.TaskMapper)
	mockMapper.On("TaskDomainToDto", dummy.TaskDomain).Return(dummy.TaskDto)
	h := NewTaskPortIn(mockTask, mockMapper)
	//Asserts
	if assert.NoError(t, h.GetAllTasks(c)) {
		err := json.Unmarshal(rec.Body.Bytes(), &data)
		if err != nil {
			log.Error(err)
		}
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, dummy.TasksDto, data)
	}

}
