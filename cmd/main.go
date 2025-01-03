package main

import (
	"log"

	todogo "github.com/FrenkenFlores/todo_go"
	"github.com/FrenkenFlores/todo_go/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	port := "8001"
	srv := new(todogo.Server)
	if err := srv.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server: %s", err.Error())
	}
}
