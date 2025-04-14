package services

import (
	"errors"

	model "github.com/greenslonik10/TODO-list/internal/models"
	repo "github.com/greenslonik10/TODO-list/internal/repositories"
)

type TaskService interface {
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task) error
	DeleteTask(id int) error
	ListTasks() ([]*model.Task, error)
}

type taskService struct {
	Repo repo.TaskRepository
}

func NewTaskService(repo repo.TaskRepository) TaskService {
	return &taskService{Repo: repo}
}

func (s *taskService) CreateTask(task *model.Task) error {
	err := s.Repo.Create(task)
	if err != nil {
		return errors.New("failed to create task")
	}
	return nil
}

func (s *taskService) UpdateTask(task *model.Task) error {
	err := s.Repo.Update(task)
	if err != nil {
		return errors.New("failed to update task")
	}

	return nil
}

func (s *taskService) DeleteTask(id int) error {
	return s.Repo.Delete(id)
}

func (s *taskService) ListTasks() ([]*model.Task, error) {
	tasks, err := s.Repo.GetAll()
	if err != nil {
		return nil, errors.New("internal server error")
	}
	if len(tasks) == 0 {
		return nil, errors.New("no tasks found")
	}
	return tasks, nil
}
