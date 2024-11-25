package repository

import (
	"database/sql"
)

func (r *Repository) GetTaskID(id string) (Task, error) {
	var task Task
	query := "SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

func (r *Repository) UpdateTaskDate(id string, newDate string) (sql.Result, error) {
	query := "UPDATE scheduler SET date = ? WHERE id = ?"
	res, err := r.db.Exec(query, newDate, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *Repository) DeleteTaskByID(id string) (sql.Result, error) {
	query := "DELETE FROM scheduler WHERE id = ?"
	res, err := r.db.Exec(query, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
