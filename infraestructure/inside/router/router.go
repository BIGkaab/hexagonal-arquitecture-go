package router

import (
	"github.com/BIGKaab/hexagonal-arquitecture-go/infraestructure/inside/controller"
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
	v1 := e.Group("/api")
	tasks := v1.Group("/tasks")
	{
		tasks.GET("", controller.GetAllTasks)
		tasks.POST("", controller.AddTask)
		tasks.GET("/:id", controller.FindTaskById)
		tasks.PUT("/:id", controller.UpdateTask)
		tasks.DELETE("/:id", controller.DeleteTask)
	}
	swagger.SwaggerRoute(e)
}
