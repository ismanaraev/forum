package main

import (
	db "forum3/db/SQLite"
	"forum3/internal/handler"
	"forum3/internal/repository"
	"forum3/internal/service"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// statment, _ := db.Database()
	db, err := db.Database()
	if err != nil {
		log.Println(err)
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	handler.InitRoutes()
}
