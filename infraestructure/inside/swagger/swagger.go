package swagger

import (
	_ "github.com/BIGKaab/hexagonal-arquitecture-go/infraestructure/inside/swagger/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func SwaggerRoute(e *echo.Echo) *echo.Echo {
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return e
}
