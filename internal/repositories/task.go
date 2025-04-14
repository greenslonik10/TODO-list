package repositories

import (
	model "github.com/greenslonik10/TODO-list/internal/models"
	"github.com/jackc/pgx"
)

type TaskRepository interface {
	Create(task *model.Task) error
	Update(task *model.Task) error
	Delete(id int) error
	GetAll() ([]*model.Task, error)
}

type taskRepository struct {
	conn *pgx.Conn
}

func NewTaskRepository(conn *pgx.Conn) TaskRepository {
	return &taskRepository{conn: conn}
}

func (r *taskRepository) Create(task *model.Task) error {
	query := `
		INSERT INTO tasks (title, description, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	err := r.conn.QueryRow(query, task.Title, task.Description, task.Status, task.CreatedAt, task.UpdatedAt).Scan(&task.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *taskRepository) Update(task *model.Task) error {
	query := `
		UPDATE tasks
		SET title = $1,
		    description = $2,
		    status = $3,
		    updated_at = $4
		WHERE id = $5
	`
	_, err := r.conn.Exec(query, task.Title, task.Description, task.Status, task.UpdatedAt, task.ID)
	return err
}

func (r *taskRepository) Delete(id int) error {
	query := `
		DELETE FROM tasks
		WHERE id = $1
	`
	_, err := r.conn.Exec(query, id)
	return err
}

func (r *taskRepository) GetAll() ([]*model.Task, error) {
	query := `
		SELECT id, title, description, status, created_at, updated_at
		FROM tasks
	`
	rows, err := r.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*model.Task
	for rows.Next() {
		task := new(model.Task)
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
