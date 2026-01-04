package todo

import "database/sql"

// Interface (what the service depends on)
type Repository interface {
	Create(title string) (Task, error)
}

// Concrete implementation (Postgres)
type PostgresRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(title string) (Task, error) {
	var task Task

	err := r.db.QueryRow(
		`INSERT INTO tasks (title)
		 VALUES ($1)
		 RETURNING id, title, done`,
		title,
	).Scan(&task.ID, &task.Title, &task.Done)

	return task, err
}
