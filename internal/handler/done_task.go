package handler

import (
	"final_project/internal/utils"
	"net/http"
	"time"
)

func (h *Handler) DoneTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, `{"error":"идентификатор задачи не указан"}`, http.StatusBadRequest)
		return
	}

	now := time.Now()

	task, err := h.repo.GetTaskID(id)
	if err != nil {
		http.Error(w, `{"error":"задача не найдена"}`, http.StatusNotFound)
		return
	}

	if task.Repeat != "" {

		newDate, err := utils.NextDate(now, task.Date, task.Repeat)
		if err != nil {
			http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
			return
		}

		_, err = h.repo.UpdateTaskDate(task.ID, newDate)
		if err != nil {
			http.Error(w, `{"error":"не удалось обновить задачу"}`, http.StatusInternalServerError)
			return
		}

	} else {

		_, err = h.repo.DeleteTaskByID(task.ID)
		if err != nil {
			http.Error(w, `{"error":"не удалось удалить задачу"}`, http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{}"))
}
