package repository

import (
	"database/sql"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment,omitempty"`
	Repeat  string `json:"repeat,omitempty"`
}

func (r *Repository) GetList() ([]Task, error) {
	var tasks []Task

	query := `SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date ASC`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *Repository) GetTaskID(id string) (Task, error) {
	var task Task
	query := "SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

func (r *Repository) DeleteTaskByID(id string) (sql.Result, error) {
	query := "DELETE FROM scheduler WHERE id = ?"
	res, err := r.db.Exec(query, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	return res, nil
}
