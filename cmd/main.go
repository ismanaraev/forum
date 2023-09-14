package main

import (
	db "forumv2/db/SQLite3"
	"forumv2/internal/handler"
	"forumv2/internal/repository"
	"forumv2/internal/service"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbPath := os.Getenv("DATABASE_PATH")
	if dbPath == "" {
		dbPath = "./db/SQLite3/store.db"
	}
	db, err := db.Database(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8081"
	}
	serverHost := os.Getenv("SERVER_ADDR")
	if serverHost == "" {
		serverHost = "0.0.0.0"
	}
	handler.InitRoutes(serverHost, serverPort)
}
