package repository

import "database/sql"

type Ingredient interface {
}

type Recipe interface {
}

type Repository struct {
	Recipe     Recipe
	Ingredient Ingredient
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{}
}
