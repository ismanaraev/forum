package main

import (
	"forum/internal/repository"
	"forum/internal/service"
)

func main() {
	db, err := repository.InitDB()
	if err != nil {
		panic(err)
	}
	repo := repository.NewRepository(db)
	service := service.NewService(repo)
}
