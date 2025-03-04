package repository

import (
	"DZ/models"
	"database/sql"
)

type Recipe interface {
	CreateRecipe(recipe models.Recipe) (int64, error)
	GetRecipe(id int64) (models.GetRecipe, error)
	GetAllRecipes() ([]models.Recipe, error)
	UpdateRecipe(id int64, recipe models.UpdateRecipe) error
	DeleteRecipe(id int64) error
}

type Ingredient interface {
	CreateIngredient(ingredient models.Ingredient) (int64, error)
	UpdateIngredient(id int64, ingredient models.Ingredient) error
	DeleteIngredient(id int64) error
}

type Repository struct {
	Recipe     Recipe
	Ingredient Ingredient
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Recipe:     NewRecipeRepository(db),
		Ingredient: NewIngredientRepository(db),
	}
}
