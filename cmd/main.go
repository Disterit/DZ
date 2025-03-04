package main

import (
	"DZ/internal/handler"
	"DZ/internal/repository"
	"DZ/internal/service"
	"DZ/pkg/logger"
	"DZ/server"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		logger.Log.Error("Error loading .env file")
		return
	}

	if err := readInConfig(); err != nil {
		logger.Log.Error("Error loading config file")
		return
	}

	db, err := repository.Connection(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.username"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DB:       viper.GetString("db.database"),
		SSL:      viper.GetString("db.sslmode"),
	})
	if err != nil {
		logger.Log.Error("Error connecting to database")
		return
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(server.HTTPServer)

	if err = srv.Run(handlers.InitRoutes(), viper.GetString("port")); err != nil {
		logger.Log.Error("Error starting server")
		return
	}

}

func readInConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
