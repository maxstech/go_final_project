package handler

import (
	"encoding/json"
	"final_project/internal/repository"
	"final_project/internal/utils"
	"log"
	"net/http"
)

func (h *Handler) AddTask(w http.ResponseWriter, r *http.Request) {
	var task repository.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "ошибка десериализации JSON", http.StatusBadRequest)
		log.Println("ошибка десериализации JSON:", err)
		return
	}

	if err := utils.ValidateTitle(task.Title); err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
		log.Println("Ошибка:", err)
		return
	}

	validatedDate, err := utils.CheckDate(task.Date, task.Repeat)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
		log.Println("ошибка валидации даты:", err)
		return
	}
	task.Date = validatedDate

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

	if _, err := w.Write(responseJSON); err != nil {
		log.Println("ошибка записи в ResponseWriter:", err)
	}
}
