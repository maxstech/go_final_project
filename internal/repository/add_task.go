package repository

import (
	"database/sql"
)

func (r *Repository) AddTask(title, date, comment, repeat string) (sql.Result, error) {

	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)`

	res, err := r.db.Exec(query, date, title, comment, repeat)
	if err != nil {
		return nil, err
	}

	return res, nil
}
