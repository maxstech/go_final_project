package handler

import (
	"encoding/json"
	"final_project/internal/utils"
	"log"
	"net/http"
	"time"
)

type Task struct {
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment,omitempty"`
	Repeat  string `json:"repeat,omitempty"`
}

func (h *Handler) AddTask(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	var task Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "ошибка десериализации JSON", http.StatusBadRequest)
		log.Println("ошибка десериализации JSON:", err)
		return
	}

	if task.Title == "" {
		http.Error(w, `{"error":"не указан заголовок задачи"}`, http.StatusBadRequest)
		log.Println("Ошибка: не указан заголовок задачи")
		return
	}
	now := time.Now()
	nowString := now.Format("20060102")

	if task.Date == "" {
		task.Date = nowString
	}
	if _, err := time.Parse("20060102", task.Date); err != nil {
		http.Error(w, `{"error":"ошибка формата даты"}`, http.StatusBadRequest)
		log.Println("ошибка формата даты:", err)
		return
	}

	if task.Date < nowString {
		if task.Repeat == "" {
			task.Date = nowString

		} else {
			nextDate, err := utils.NextDate(now, task.Date, task.Repeat)
			if err != nil {
				http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
				log.Println("ошибка вычисления следующей даты:", err)
				return
			}
			task.Date = nextDate
		}
	}

	res, err := h.repo.AddTask(task.Title, task.Date, task.Comment, task.Repeat)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		log.Println("ошибка добавления задачи в репозиторий:", err)
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		http.Error(w, `{"error":"ошибка получения ID задачи"}`, http.StatusInternalServerError)
		log.Println("ошибка получения ID задачи:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	responseData := map[string]interface{}{"id": id}
	responseJSON, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, `{"error":"ошибка формирования ответа"}`, http.StatusInternalServerError)
		log.Println("ошибка формирования ответа:", err)
		return
	}
	w.Write([]byte(responseJSON))

}
