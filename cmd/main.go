package main

import (
	"log"

	todogo "github.com/FrenkenFlores/todo_go"
	"github.com/FrenkenFlores/todo_go/pkg/handler"
	"github.com/FrenkenFlores/todo_go/pkg/repository"
	"github.com/FrenkenFlores/todo_go/pkg/service"
	"github.com/spf13/viper"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error occured durring initialization: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	port := viper.GetString("port")
	srv := new(todogo.Server)
	if err := srv.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server: %s", err.Error())
	}
}
