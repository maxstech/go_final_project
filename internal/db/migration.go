package db

import (
	"final_project/internal/repository"
	"log"
	"os"
	"path/filepath"
)

func Migration(rep *repository.Repository) {

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
