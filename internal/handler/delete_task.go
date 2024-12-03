package handler

import (
	"net/http"
)

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		http.Error(w, "метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, `{"error":"идентификатор задачи не указан"}`, http.StatusBadRequest)
		return
	}

	task, err := h.repo.GetTaskID(id)
	if err != nil {
		http.Error(w, `{"error":"задача не найдена"}`, http.StatusNotFound)
		return
	}

	_, err = h.repo.DeleteTaskByID(task.ID)
	if err != nil {
		http.Error(w, `{"error":"не удалось удалить задачу"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{}"))
}
