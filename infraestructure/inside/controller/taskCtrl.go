package controller

import (
	"clean-arquitecture-go/application/mapper"
	"clean-arquitecture-go/application/mapper/impl"
	"clean-arquitecture-go/application/port/in"
	"clean-arquitecture-go/application/usescases"
	"clean-arquitecture-go/infraestructure/inside/dto"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
)

var taskPortIn in.TaskPortIn = usescases.NewTaskPortOut()
var mappers mapper.TaskMapper = impl.NewTaskMapperImpl()

func NewTaskPortIn(portIn in.TaskPortIn, mapper mapper.TaskMapper) {
	taskPortIn = portIn
	mappers = mapper
}

// GetAllTasks dogoc
// @Tags Tasks
// @Summary Get all tasks
// @Description return all tasks
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.Message
// @Failure 400 {object} dto.MessageError
// @Failure 500 {object} dto.MessageError
// @Router /tasks [get]
func GetAllTasks(c echo.Context) error {
	var tasksDto []dto.Task
	tasksDomain, err := taskPortIn.InGetAllTasks()
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	for _, taskDomain := range tasksDomain {
		tasksDto = append(tasksDto, mappers.TaskDomainToDto(taskDomain))
	}
	return c.JSON(http.StatusOK, dto.Message{
		Message: fmt.Sprintf("enums.MessageSuccessfully, enums.MessageTasks, enums.MessageLoaded"),
		Data:    tasksDto,
	})
}

// AddTask dogoc
// @Tags Tasks
// @Summary Add task
// @Description return new tasks
// @Accept  json
// @Produce  json
// @Param task body	dto.Task true "Task"
// @Success 200 {object} dto.Message
// @Failure 400 {object} dto.MessageError
// @Failure 500 {object} dto.MessageError
// @Router /tasks [post]
func AddTask(c echo.Context) error {
	var data dto.Task

	if err := c.Bind(&data); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := data.Validate(); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	dataDomain := mappers.TaskDtoToDomain(data)

	res, err := taskPortIn.InAddTask(dataDomain)

	dataDto := mappers.TaskDomainToDto(res)

	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, dto.Message{
		Message: "CREATE",
		Data:    dataDto,
	})
}

// FindTaskById dogoc
// @Tags Tasks
// @Summary Find Task By ID
// @Description return Find Task By ID
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} dto.Message
// @Failure 400 {object} dto.MessageError
// @Failure 500 {object} dto.MessageError
// @Router /tasks/{id} [get]
func FindTaskById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	dataDomain, err := taskPortIn.InFindTaskById(id)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	dataDto := mappers.TaskDomainToDto(dataDomain)

	return c.JSON(http.StatusOK, dto.Message{
		Message: "FIND",
		Data:    dataDto,
	})
}

// UpdateTask dogoc
// @Tags Tasks
// @Summary Update Task By ID
// @Description return Update Task By ID
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param task body	dto.Task true "Task"
// @Success 200 {object} dto.Message
// @Failure 400 {object} dto.MessageError
// @Failure 500 {object} dto.MessageError
// @Router /tasks/{id} [put]
func UpdateTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var data dto.Task

	if err := c.Bind(&data); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := data.Validate(); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	dataDomain := mappers.TaskDtoToDomain(data)

	res, err := taskPortIn.InUpdateTask(id, dataDomain)

	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	dataDto := mappers.TaskDomainToDto(res)

	return c.JSON(http.StatusCreated, dto.Message{
		Message: "UPDATE",
		Data:    dataDto,
	})
}

// DeleteTask dogoc
// @Tags Tasks
// @Summary Deleted Task By ID
// @Description return void
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} dto.Message
// @Failure 400 {object} dto.MessageError
// @Failure 500 {object} dto.MessageError
// @Router /tasks/{id} [delete]
func DeleteTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := taskPortIn.InDeleteTask(id)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, dto.Message{
		Message: fmt.Sprintf("enums.MessageSuccessfully, enums.MessageTask, enums.MessageDeleted"),
	})
}
