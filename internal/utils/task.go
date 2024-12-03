package utils

import (
	"errors"
	"time"
)

func ValidateTitle(title string) error {
	if title == "" {
		return errors.New("не указан заголовок задачи")
	}
	return nil
}

func CheckDate(date, repeat string) (string, error) {
	now := time.Now()
	nowString := now.Format(DateFormat)

	if date == "" {
		return nowString, nil
	}
	if _, err := time.Parse(DateFormat, date); err != nil {
		return "", errors.New("ошибка формата даты")
	}
	if date < nowString {
		if repeat == "" {
			return nowString, nil
		} else {
			nextDate, err := NextDate(now, date, repeat)
			if err != nil {
				return "", err
			}
			return nextDate, nil
		}
	}
	return date, nil
}
