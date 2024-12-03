package handler

import (
	"encoding/json"
	"final_project/internal/repository"
	"fmt"
	"log"
	"net/http"
)

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, `{"error":"не указан идентификатор"}`, http.StatusBadRequest)
		return
	}

	tasks, err := h.repo.GetList()
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		log.Println("ошибка получения списка задач:", err)
		return
	}

	var foundTask *repository.Task
	for _, task := range tasks {
		if task.ID == id {
			foundTask = &task
			break
		}
	}

	if foundTask == nil {
		http.Error(w, `{"error":"задача не найдена"}`, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(foundTask); err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		log.Println("ошибка сериализации JSON:", err)
	}
}

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task repository.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusBadRequest)
		return
	}

	log.Printf("Receive request to update task: ID=%s", task.ID)

	if task.Title == "" {
		http.Error(w, `{"error":"заголовок не может быть пустым"}`, http.StatusBadRequest)
		return
	}

	err := h.repo.UpdateTask(task)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusNotFound)
		log.Println("ошибка обновления задачи:", err)
		return
	}

	log.Printf("Task updated successfully: ID=%s", task.ID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}
