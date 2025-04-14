package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/greenslonik10/TODO-list/internal/controllers"
	"github.com/greenslonik10/TODO-list/internal/database/postgresql"
	"github.com/greenslonik10/TODO-list/internal/repositories"
	"github.com/greenslonik10/TODO-list/internal/routes"
	"github.com/greenslonik10/TODO-list/internal/services"
)

func main() {
	conn, err := postgresql.NewPostgresConn()
	if err != nil {
		panic(err)
	}

	repo := repositories.NewTaskRepository(conn)
	service := services.NewTaskService(repo)
	controller := controllers.NewTaskController(service)
	route := routes.NewAuthRoutes(controller)

	app := fiber.New()

	route.InitRoutes(app)

	app.Listen(":3000")
}
