package main

import (
	"final_project/internal/db"
	"final_project/internal/handler"
	"final_project/internal/repository"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {

	dbase := db.New()
	defer dbase.Close()

	repo := repository.New(dbase)

	db.Migration(repo)

	handlerAddTask := handler.New(repo)
	handlerNextDate := handler.New(repo)
	handlerGetTasks := handler.New(repo)
	handlerGetTask := handler.New(repo)
	handlerUpdateTask := handler.New(repo)
	handlerDoneTask := handler.New(repo)
	handlerDeleteTask := handler.New(repo)

	r := chi.NewRouter()
	r.Handle("/*", http.FileServer(http.Dir("./web/")))
	r.Post("/api/task", handlerAddTask.AddTask)
	r.Get("/api/nextdate", handlerNextDate.NextDateHandler)
	r.Get("/api/tasks", handlerGetTasks.GetTasks)
	r.Delete("/api/task", handlerDeleteTask.DeleteTask)
	r.Get("/api/task", handlerGetTask.GetTask)
	r.Put("/api/task", handlerUpdateTask.UpdateTask)
	r.Post("/api/task/done", handlerDoneTask.DoneTask)

	port := ":7540"
	fmt.Printf("Сервер запущен на порту%s\n", port)

	if err := http.ListenAndServe(port, r); err != nil {
		fmt.Println("Ошибка запуска сервера", err)
	}
}
