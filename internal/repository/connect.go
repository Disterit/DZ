package repository

import (
	"DZ/pkg/logger"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	recipeTable     = "recipe"
	ingredientTable = "ingredient"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DB       string
	SSL      string
}

func Connection(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DB, cfg.SSL))
	if err != nil {
		logger.Log.Error("error connecting to the database", err)
		return nil, err
	}

	return db, err
}
