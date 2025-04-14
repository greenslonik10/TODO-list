package controllers

import (
	"strconv"

	model "github.com/greenslonik10/TODO-list/internal/models"
	service "github.com/greenslonik10/TODO-list/internal/services"

	"github.com/gofiber/fiber/v2"
)

type TaskController interface {
	CreateTask(ctx *fiber.Ctx) error
	UpdateTask(ctx *fiber.Ctx) error
	DeleteTask(ctx *fiber.Ctx) error
	ListTasks(ctx *fiber.Ctx) error
}

type taskController struct {
	TaskService service.TaskService
}

func NewTaskController(taskService service.TaskService) TaskController {
	return &taskController{TaskService: taskService}
}

func (c taskController) CreateTask(ctx *fiber.Ctx) error {

	var task model.Task

	if err := ctx.BodyParser(&task); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := c.TaskService.CreateTask(&task)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}

func (c taskController) UpdateTask(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid task ID",
		})
	}

	var task model.Task

	if err := ctx.BodyParser(&task); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	task.ID = id

	err = c.TaskService.UpdateTask(&task)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})

}

func (c taskController) DeleteTask(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid task ID",
		})
	}

	err = c.TaskService.DeleteTask(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error deleting task",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}

func (c *taskController) ListTasks(ctx *fiber.Ctx) error {

	tasks, err := c.TaskService.ListTasks()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if len(tasks) == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "no tasks found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"tasks":   tasks,
	})

}
