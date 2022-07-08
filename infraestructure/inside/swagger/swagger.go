package swagger

import (
	"github.com/BIGKaab/hexagonal-arquitecture-go/infraestructure/inside/enum"
	_ "github.com/BIGKaab/hexagonal-arquitecture-go/infraestructure/inside/swagger/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func SwaggerRoute(e *echo.Echo) *echo.Echo {
	e.GET(enum.ROUTER_SWAGGER, echoSwagger.WrapHandler)

	return e
}
