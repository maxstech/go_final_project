package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func NextDate(now time.Time, date string, repeat string) (string, error) {

	if repeat == "" {
		return "", nil
	}

	if repeat == "w" {
		return "", errors.New("правило повторения не реализуется")
	}
	if repeat == "m" {
		return "", errors.New("правило повторения не реализуется")
	}

	startDate, err := time.Parse("20060102", date)
	if err != nil {

		return "", fmt.Errorf("некорректная дата: %v", err)
	}

	nextDate := startDate

	switch {
	case strings.HasPrefix(repeat, "d"):

		daysStr := strings.TrimSpace(strings.TrimPrefix(repeat, "d"))

		days, err := strconv.Atoi(daysStr)
		if err != nil || days <= 0 || days > 400 {

			return "", errors.New("недопустимое значение для правила d")
		}

		nextDate = nextDate.AddDate(0, 0, days)
		for nextDate.Before(now) || nextDate.Equal(now) {

			nextDate = nextDate.AddDate(0, 0, days)
		}

	case repeat == "y":

		nextDate = nextDate.AddDate(1, 0, 0)

		for nextDate.Before(now) || nextDate.Equal(now) {

			nextDate = nextDate.AddDate(1, 0, 0)
		}

	default:

		return "", errors.New("неподдерживаемый формат правила повторения")
	}

	return nextDate.Format("20060102"), nil
}
