package controller

import (
	"github.com/BIGKaab/hexagonal-arquitecture-go/application/mapper"
	"github.com/BIGKaab/hexagonal-arquitecture-go/application/port/in"
	"github.com/BIGKaab/hexagonal-arquitecture-go/infraestructure/inside/dto"
	"github.com/BIGKaab/hexagonal-arquitecture-go/infraestructure/inside/enum"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
)

type controller struct {
	portIn in.TaskPortIn
	mapper mapper.TaskMapper
}

func NewTaskPortIn(portIn in.TaskPortIn, mapper mapper.TaskMapper) *controller {
	return &controller{portIn, mapper}
}

// GetAllTasks dogoc
// @Tags Tasks
// @Summary Get all tasks
// @Description return all tasks
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.Task
// @Failure 400 {object} dto.MessageError
// @Failure 500 {object} dto.MessageError
// @Router /tasks [get]
func (ctrl *controller) GetAllTasks(c echo.Context) error {
	var tasksDto []dto.Task
	tasksDomain, err := ctrl.portIn.InGetAllTasks()
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	for _, taskDomain := range tasksDomain {
		tasksDto = append(tasksDto, ctrl.mapper.TaskDomainToDto(taskDomain))
	}
	return c.JSON(http.StatusOK, tasksDto)

}

// AddTask dogoc
// @Tags Tasks
// @Summary Add task
// @Description return new tasks
// @Accept  json
// @Produce  json
// @Param task body	dto.Task true "Task"
// @Success 200 {object} dto.Task
// @Failure 400 {object} dto.MessageError
// @Failure 500 {object} dto.MessageError
// @Router /tasks [post]
func (ctrl *controller) AddTask(c echo.Context) error {
	var data dto.Task

	if err := c.Bind(&data); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := data.Validate(); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	dataDomain := ctrl.mapper.TaskDtoToDomain(data)

	res, err := ctrl.portIn.InAddTask(dataDomain)

	dataDto := ctrl.mapper.TaskDomainToDto(res)

	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, dataDto)
}

// FindTaskById dogoc
// @Tags Tasks
// @Summary Find Task By ID
// @Description return Find Task By ID
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} dto.Task
// @Failure 400 {object} dto.MessageError
// @Failure 404 {object} dto.MessageError
// @Failure 500 {object} dto.MessageError
// @Router /tasks/{id} [get]
func (ctrl *controller) FindTaskById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param(enum.ROUTER_ID))
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	dataDomain, err := ctrl.portIn.InFindTaskById(id)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	dataDto := ctrl.mapper.TaskDomainToDto(dataDomain)

	return c.JSON(http.StatusOK, dataDto)
}

// UpdateTask dogoc
// @Tags Tasks
// @Summary Update Task By ID
// @Description return Update Task By ID
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param task body	dto.Task true "Task"
// @Success 200 {object} dto.Task
// @Failure 400 {object} dto.MessageError
// @Failure 404 {object} dto.MessageError
// @Failure 500 {object} dto.MessageError
// @Router /tasks/{id} [put]
func (ctrl *controller) UpdateTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param(enum.ROUTER_ID))
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	var data dto.Task

	if err := c.Bind(&data); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := data.Validate(); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	dataDomain := ctrl.mapper.TaskDtoToDomain(data)

	res, err := ctrl.portIn.InUpdateTask(id, dataDomain)

	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	dataDto := ctrl.mapper.TaskDomainToDto(res)

	return c.JSON(http.StatusOK, dataDto)
}

// DeleteTask dogoc
// @Tags Tasks
// @Summary Deleted Task By ID
// @Description return void
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 204
// @Failure 400 {object} dto.MessageError
// @Failure 500 {object} dto.MessageError
// @Router /tasks/{id} [delete]
func (ctrl *controller) DeleteTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param(enum.ROUTER_ID))
	err := ctrl.portIn.InDeleteTask(id)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusNoContent, nil)
}
