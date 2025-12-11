package tasks

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SQLRepository struct {
	db *pgxpool.Pool
}

func NewSQLRepository(db *pgxpool.Pool) *SQLRepository {
	return &SQLRepository{db: db}
}

// LIST
func (r *SQLRepository) List() ([]Task, error) {
	rows, err := r.db.Query(context.Background(),
		"SELECT id, title, done, created_at, completed_at FROM tasks ORDER BY id",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var t Task
		err := rows.Scan(&t.ID, &t.Title, &t.Done, &t.CreatedAt, &t.CompletedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

// ADD
func (r *SQLRepository) Add(title string) (int, error) {
	var id int
	err := r.db.QueryRow(context.Background(),
		"INSERT INTO tasks (title, done, created_at) VALUES ($1, false, NOW()) RETURNING id",
		title,
	).Scan(&id)

	return id, err
}

// COMPLETE
func (r *SQLRepository) Complete(id int) error {
	res, err := r.db.Exec(context.Background(),
		"UPDATE tasks SET done = true, completed_at = NOW() WHERE id = $1",
		id,
	)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return errors.New("task inesistente")
	}

	return nil
}

// DELETE
func (r *SQLRepository) Delete(id int) error {
	res, err := r.db.Exec(context.Background(),
		"DELETE FROM tasks WHERE id = $1",
		id,
	)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return errors.New("task inesistente")
	}

	return nil
}
