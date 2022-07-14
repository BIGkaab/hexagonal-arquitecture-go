package router

import (
	"github.com/BIGKaab/hexagonal-arquitecture-go/application/mapper/impl"
	"github.com/BIGKaab/hexagonal-arquitecture-go/application/usescases"
	"github.com/BIGKaab/hexagonal-arquitecture-go/infraestructure/inside/controller"
	"github.com/BIGKaab/hexagonal-arquitecture-go/infraestructure/inside/enum"
	"github.com/BIGKaab/hexagonal-arquitecture-go/infraestructure/inside/swagger"

	"github.com/labstack/echo/v4"
)

// @title Tasks API
// @version 1.0
// @description Tasks Manager.
// @termsOfService http://swagger.io/terms/
// @contact.name Kiven Acevedo
// @contact.email kiven.acv2@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3005
// @BasePath /api
func Routes(e *echo.Echo) {
	v1 := e.Group(enum.ROUTER_GROUP_GLOBAL)
	tasks := v1.Group(enum.ROUTER_TASK_GROUP)
	t := controller.NewTaskPortIn(usescases.NewTaskPortOut(), impl.NewTaskMapperImpl())
	{
		tasks.GET("", t.GetAllTasks)
		tasks.POST("", t.AddTask)
		tasks.GET(enum.ROUTER_PARAM_ID, t.FindTaskById)
		tasks.PUT(enum.ROUTER_PARAM_ID, t.UpdateTask)
		tasks.DELETE(enum.ROUTER_PARAM_ID, t.DeleteTask)
	}
	swagger.SwaggerRoute(e)
}
