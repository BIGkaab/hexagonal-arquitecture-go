package main

import (
	"clean-arquitecture-go/infraestructure/inside/router"
	"clean-arquitecture-go/infraestructure/outside/gorm/config"
	"clean-arquitecture-go/infraestructure/outside/gorm/migration"
	"clean-arquitecture-go/infraestructure/outside/gorm/seeder"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo initialization
	e := echo.New()

	//Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	router.Routes(e)

	// DB connection
	config.ConnInstance()
	// DB Migrations
	migration.Execute()
	// DB Seeders
	seeder.Execute()

	// Start the server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":3005")))
}
