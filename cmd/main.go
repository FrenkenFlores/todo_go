package main

import (
	"log"

	todogo "github.com/FrenkenFlores/todo_go"
	"github.com/FrenkenFlores/todo_go/pkg/handler"
	"github.com/FrenkenFlores/todo_go/pkg/repository"
	"github.com/FrenkenFlores/todo_go/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	port := "8001"
	srv := new(todogo.Server)
	if err := srv.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server: %s", err.Error())
	}
}
