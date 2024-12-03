package repository

import (
	"database/sql"
	"errors"
	"final_project/internal/utils"
	"regexp"
	"time"
)

func isValidDate(date string) bool {
	match, _ := regexp.MatchString(`^\d{8}$`, date)
	if !match {
		return false
	}

	_, err := time.Parse(utils.DateFormat, date)
	return err == nil
}

func isValidRepeat(repeat string) bool {
	return repeat == "" || regexp.MustCompile(`^(d \d+|y)$`).MatchString(repeat)
}

func (r *Repository) GetTaskByID(id string) (Task, error) {
	var task Task
	query := "SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?"
	if err := r.db.QueryRow(query, id).Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
		if err == sql.ErrNoRows {
			return Task{}, errors.New("задача не найдена")
		}
		return Task{}, err
	}
	return task, nil
}

func (r *Repository) UpdateTask(task Task) error {
	if !isValidDate(task.Date) {
		return errors.New("некорректный формат даты; ожидается YYYYMMDD")
	}

	if !isValidRepeat(task.Repeat) {
		return errors.New("некорректный формат для repeat")
	}

	query := "UPDATE scheduler SET date = ?, title = ?, comment = ?, repeat = ? WHERE id = ?"
	result, err := r.db.Exec(query, task.Date, task.Title, task.Comment, task.Repeat, task.ID)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRows == 0 {
		return errors.New("задача не найдена для обновления")
	}

	return nil
}
