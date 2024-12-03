package handler

import (
	"encoding/json"
	"final_project/internal/repository"
	"fmt"
	"log"
	"net/http"
)

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.repo.GetList()
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		log.Println("ошибка получения списка задач:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if len(tasks) == 0 {
		if err := json.NewEncoder(w).Encode(map[string][]repository.Task{"tasks": {}}); err != nil {
			http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
			log.Println("ошибка сериализации JSON:", err)
		}
		return
	}

	if err := json.NewEncoder(w).Encode(map[string][]repository.Task{"tasks": tasks}); err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		log.Println("ошибка сериализации JSON:", err)
	}
}
