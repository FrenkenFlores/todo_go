package main

import (
	"log"

	todogo "github.com/FrenkenFlores/todo_go"
)

func main() {
	port := "8001"
	srv := new(todogo.Server)
	if err := srv.Run(port); err != nil {
		log.Fatalf("error occured while running server: %s", err.Error())
	}
}
