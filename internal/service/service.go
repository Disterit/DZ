package service

import (
	"DZ/internal/repository"
	"DZ/models"
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

type Service struct {
	Recipe     Recipe
	Ingredient Ingredient
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Recipe:     NewRecipeService(repo.Recipe),
		Ingredient: NewIngredientService(repo.Ingredient),
	}
}
