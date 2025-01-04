package main

import (
	"log"
	"os"

	todogo "github.com/FrenkenFlores/todo_go"
	"github.com/FrenkenFlores/todo_go/pkg/handler"
	"github.com/FrenkenFlores/todo_go/pkg/repository"
	"github.com/FrenkenFlores/todo_go/pkg/service"
	"github.com/joho/godotenv"
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
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error occured while loading env: %s", err.Error())
	}
	db, err := repository.NewDatabase(repository.Config{
		Host:     viper.GetStringMapString("db")["host"],
		Port:     viper.GetStringMapString("db")["port"],
		User:     viper.GetStringMapString("db")["user"],
		Password: os.Getenv("DB_PASSWORD"),
		Database: viper.GetStringMapString("db")["database"],
		SslMode:  viper.GetStringMapString("db")["sslmode"],
	})
	if err != nil {
		log.Fatalf("failed to initialize database: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	port := viper.GetString("port")
	srv := new(todogo.Server)
	if err := srv.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server: %s", err.Error())
	}
}
