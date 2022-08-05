package main

import (
	"fmt"
	"github.com/BIGKaab/hexagonal-arquitecture-go/infraestructure/inside/router"
	"github.com/BIGKaab/hexagonal-arquitecture-go/infraestructure/outside/gorm/config"
	"github.com/BIGKaab/hexagonal-arquitecture-go/infraestructure/outside/gorm/migration"
	"github.com/BIGKaab/hexagonal-arquitecture-go/infraestructure/outside/gorm/seeder"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/subosito/gotenv"
	"os"
)

func main() {
	//Environment
	gotenv.Load()

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
	//e.Logger.Fatal(e.Start(fmt.Sprintf(":3005")))
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("API_PORT"))))
}
