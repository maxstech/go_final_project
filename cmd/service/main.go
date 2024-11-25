package main

import (
	"final_project/internal/db"
	"final_project/internal/handler"
	"final_project/internal/repository"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
)

func main() {

	db := db.New()
	repo := repository.New(db)
	migration(repo)

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
	r.Delete("/api/tasks", handlerGetTasks.GetTasks)
	r.Delete("/api/task", handlerDeleteTask.DeleteTask)
	r.Get("/api/task", handlerGetTask.GetTask)
	r.Put("/api/task", handlerUpdateTask.UpdateTask)
	r.Post("/api/task/done", handlerDoneTask.DoneTask)

	if err := http.ListenAndServe(":7540", r); err != nil {
		fmt.Println("Ошибка запуска сервера", err)

	}

}

func migration(rep *repository.Repository) {

	appPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dbFile := filepath.Join(filepath.Dir(appPath), "scheduler.db")
	_, err = os.Stat(dbFile)

	var install bool
	if err != nil {
		install = true
	}
	if install {
		if err := rep.CreateScheduler(); err != nil {
			log.Fatal(err)
		}
	}
}
