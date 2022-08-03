package controller

import (
	"github.com/BIGKaab/hexagonal-arquitecture-go/domain"
	"github.com/BIGKaab/hexagonal-arquitecture-go/dummy"
	"github.com/BIGKaab/hexagonal-arquitecture-go/infraestructure/inside/enum"
	"github.com/BIGKaab/hexagonal-arquitecture-go/mocks"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetAllTasks(t *testing.T) {
	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		ExpectedData       []domain.Task
		ExpectedDataJson   string
		ExpectedError      error
	}{
		{
			Name:               "OK",
			ExpectedStatusCode: http.StatusOK,
			ExpectedData:       dummy.TasksDomain,
			ExpectedError:      nil,
		},
		{
			Name:               "Internal server error",
			ExpectedStatusCode: http.StatusInternalServerError,
			ExpectedData:       nil,
			ExpectedError:      enum.INTERNAL_SERVER_ERROR,
		},
	}

	mapper := new(mocks.TaskMapper)
	mapper.On("TaskDomainToDto", dummy.TaskDomain).Return(dummy.TaskDto)

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			//Setup
			e := echo.New()
			r := httptest.NewRequest(http.MethodGet, enum.ROUTER_GROUP_GLOBAL, nil)
			w := httptest.NewRecorder()

			c := e.NewContext(r, w)

			service := new(mocks.TaskPortIn)
			service.On("InGetAllTasks").Return(tc.ExpectedData, tc.ExpectedError)
			h := NewTaskPortIn(service, mapper)

			//Asserts
			err := h.GetAllTasks(c)
			if err != nil {
				t.Errorf("unexpected error := %s", err)
			}

			if w.Code != tc.ExpectedStatusCode {
				t.Errorf("expected status code %d, got %d", tc.ExpectedStatusCode, w.Code)
			}
		})
	}
}

func TestAddTask(t *testing.T) {
	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		ExpectedData       domain.Task
		InputData          string
		ExpectedError      error
	}{
		{
			Name:               "OK",
			ExpectedStatusCode: http.StatusCreated,
			ExpectedData:       dummy.TaskDomain,
			ExpectedError:      nil,
			InputData:          dummy.InputTaskJson,
		},
		{
			Name:               "Internal server error",
			ExpectedStatusCode: http.StatusInternalServerError,
			ExpectedData:       dummy.TaskDomain,
			ExpectedError:      enum.INTERNAL_SERVER_ERROR,
			InputData:          dummy.InputTaskJson,
		},
		{
			Name:               "without body",
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedData:       dummy.TaskDomain,
			ExpectedError:      enum.BAB_REQUEST,
			InputData:          "",
		},
		{
			Name:               "invalid body",
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedData:       dummy.TaskDomain,
			ExpectedError:      enum.BAB_REQUEST,
			InputData:          dummy.InputTaskJsonFailValidate,
		},
	}

	mapper := new(mocks.TaskMapper)
	mapper.On("TaskDtoToDomain", dummy.TaskDto).Return(dummy.TaskDomain)
	mapper.On("TaskDomainToDto", dummy.TaskDomain).Return(dummy.TaskDto)

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			//Setup
			e := echo.New()
			r := httptest.NewRequest(http.MethodPost, enum.ROUTER_GROUP_GLOBAL, strings.NewReader(tc.InputData))
			r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			w := httptest.NewRecorder()

			c := e.NewContext(r, w)

			service := new(mocks.TaskPortIn)
			service.On("InAddTask", dummy.TaskDomain).Return(tc.ExpectedData, tc.ExpectedError)
			h := NewTaskPortIn(service, mapper)

			//Asserts
			err := h.AddTask(c)
			if err != nil {
				t.Errorf("unexpected error := %s", err)
			}

			if w.Code != tc.ExpectedStatusCode {
				t.Errorf("expected status code %d, got %d", tc.ExpectedStatusCode, w.Code)
			}
		})
	}
}

func TestFindTaskById(t *testing.T) {
	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		ExpectedData       domain.Task
		ExpectedError      error
		IdParamString      string
		IdParamInt         int
	}{
		{
			Name:               "OK",
			ExpectedStatusCode: http.StatusOK,
			ExpectedData:       dummy.TaskDomain,
			ExpectedError:      nil,
			IdParamString:      "1",
			IdParamInt:         1,
		},
		{
			Name:               "Internal server error",
			ExpectedStatusCode: http.StatusInternalServerError,
			ExpectedData:       dummy.TaskDomain,
			ExpectedError:      enum.INTERNAL_SERVER_ERROR,
			IdParamString:      "1",
			IdParamInt:         1,
		},
		{
			Name:               "Not found",
			ExpectedStatusCode: http.StatusNotFound,
			ExpectedData:       dummy.TaskDomain,
			ExpectedError:      enum.NOT_FOUND,
			IdParamString:      "fail",
			IdParamInt:         1,
		},
	}

	mapper := new(mocks.TaskMapper)
	mapper.On("TaskDomainToDto", dummy.TaskDomain).Return(dummy.TaskDto)

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			//Setup
			e := echo.New()
			r := httptest.NewRequest(http.MethodGet, enum.ROUTER_GROUP_GLOBAL, nil)
			w := httptest.NewRecorder()
			c := e.NewContext(r, w)
			c.SetPath(enum.ROUTER_TASK_GROUP + enum.ROUTER_PARAM_ID)
			c.SetParamNames(enum.ROUTER_ID)
			c.SetParamValues(tc.IdParamString)

			service := new(mocks.TaskPortIn)
			service.On("InFindTaskById", tc.IdParamInt).Return(tc.ExpectedData, tc.ExpectedError)
			h := NewTaskPortIn(service, mapper)

			//Asserts
			err := h.FindTaskById(c)
			if err != nil {
				t.Errorf("unexpected error := %s", err)
			}

			if w.Code != tc.ExpectedStatusCode {
				t.Errorf("expected status code %d, got %d", tc.ExpectedStatusCode, w.Code)
			}
		})
	}
}

func TestUpdateTask(t *testing.T) {
	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		ExpectedData       domain.Task
		ExpectedError      error
		InputData          string
		IdParamString      string
		IdParamInt         int
	}{
		{
			Name:               "OK",
			ExpectedStatusCode: http.StatusOK,
			ExpectedData:       dummy.TaskDomain,
			ExpectedError:      nil,
			InputData:          dummy.InputTaskJson,
			IdParamString:      "1",
			IdParamInt:         1,
		},
		{
			Name:               "Internal server error",
			ExpectedStatusCode: http.StatusInternalServerError,
			ExpectedData:       dummy.TaskDomain,
			ExpectedError:      enum.INTERNAL_SERVER_ERROR,
			InputData:          dummy.InputTaskJson,
			IdParamString:      "1",
			IdParamInt:         1,
		},
		{
			Name:               "Not found",
			ExpectedStatusCode: http.StatusNotFound,
			ExpectedData:       dummy.TaskDomain,
			ExpectedError:      enum.NOT_FOUND,
			InputData:          dummy.InputTaskJson,
			IdParamString:      "fail",
			IdParamInt:         1,
		},
		{
			Name:               "Not found",
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedData:       dummy.TaskDomain,
			ExpectedError:      enum.BAB_REQUEST,
			InputData:          "",
			IdParamString:      "1",
			IdParamInt:         1,
		},
		{
			Name:               "invalid json",
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedData:       dummy.TaskDomain,
			ExpectedError:      enum.BAB_REQUEST,
			InputData:          dummy.InputTaskJsonFailValidate,
			IdParamString:      "1",
			IdParamInt:         1,
		},
	}

	mapper := new(mocks.TaskMapper)
	mapper.On("TaskDtoToDomain", dummy.TaskDto).Return(dummy.TaskDomain)
	mapper.On("TaskDomainToDto", dummy.TaskDomain).Return(dummy.TaskDto)

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			//Setup
			e := echo.New()
			r := httptest.NewRequest(http.MethodPut, enum.ROUTER_GROUP_GLOBAL, strings.NewReader(tc.InputData))
			r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			w := httptest.NewRecorder()
			c := e.NewContext(r, w)
			c.SetPath(enum.ROUTER_TASK_GROUP + enum.ROUTER_PARAM_ID)
			c.SetParamNames(enum.ROUTER_ID)
			c.SetParamValues(tc.IdParamString)

			service := new(mocks.TaskPortIn)
			service.On("InUpdateTask", tc.IdParamInt, dummy.TaskDomain).Return(tc.ExpectedData, tc.ExpectedError)
			h := NewTaskPortIn(service, mapper)

			//Asserts
			err := h.UpdateTask(c)
			if err != nil {
				t.Errorf("unexpected error := %s", err)
			}

			if w.Code != tc.ExpectedStatusCode {
				t.Errorf("expected status code %d, got %d", tc.ExpectedStatusCode, w.Code)
			}
		})
	}
}

func TestDeleteTask(t *testing.T) {
	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		ExpectedError      error
		IdParamString      string
		IdParamInt         int
	}{
		{
			Name:               "OK",
			ExpectedStatusCode: http.StatusNoContent,
			ExpectedError:      nil,
			IdParamString:      "1",
			IdParamInt:         1,
		},
		{
			Name:               "Internal server error",
			ExpectedStatusCode: http.StatusInternalServerError,
			ExpectedError:      enum.INTERNAL_SERVER_ERROR,
			IdParamString:      "1",
			IdParamInt:         1,
		},
		{
			Name:               "Not found",
			ExpectedStatusCode: http.StatusNotFound,
			ExpectedError:      enum.NOT_FOUND,
			IdParamString:      "fail",
			IdParamInt:         1,
		},
	}

	mapper := new(mocks.TaskMapper)
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			//Setup
			e := echo.New()
			r := httptest.NewRequest(http.MethodDelete, enum.ROUTER_GROUP_GLOBAL, nil)
			w := httptest.NewRecorder()
			c := e.NewContext(r, w)
			c.SetPath(enum.ROUTER_TASK_GROUP + enum.ROUTER_PARAM_ID)
			c.SetParamNames(enum.ROUTER_ID)
			c.SetParamValues(tc.IdParamString)

			service := new(mocks.TaskPortIn)
			service.On("InDeleteTask", tc.IdParamInt).Return(tc.ExpectedError)
			h := NewTaskPortIn(service, mapper)

			//Asserts
			err := h.DeleteTask(c)
			if err != nil {
				t.Errorf("unexpected error := %s", err)
			}

			if w.Code != tc.ExpectedStatusCode {
				t.Errorf("expected status code %d, got %d", tc.ExpectedStatusCode, w.Code)
			}
		})
	}
}

/*
var (
	tasks    []dto.Task
	task     dto.Task
	IdString = "1"
	Id       = 1
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
*/
