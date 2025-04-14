package routes

import (
	"github.com/gofiber/fiber/v2"
	cont "github.com/greenslonik10/TODO-list/internal/controllers"
)

type TaskRoutes interface {
	InitRoutes(a *fiber.App)
}

type taskRoutes struct {
	taskController cont.TaskController
}

func NewAuthRoutes(taskController cont.TaskController) TaskRoutes {
	return &taskRoutes{taskController: taskController}
}

func (r *taskRoutes) InitRoutes(a *fiber.App) {
	taskGroup := a.Group("/tasks")
	{
		taskGroup.Post("/", r.taskController.CreateTask)
		taskGroup.Get("/", r.taskController.ListTasks)
		taskGroup.Put("/:id", r.taskController.UpdateTask)
		taskGroup.Delete("/:id", r.taskController.DeleteTask)
	}
}
