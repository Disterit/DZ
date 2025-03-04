package service

import (
	"DZ/internal/repository"
	"DZ/models"
)

type RecipeService struct {
	repo repository.Recipe
}

func NewRecipeService(repo repository.Recipe) *RecipeService {
	return &RecipeService{repo: repo}
}

func (s *RecipeService) CreateRecipe(recipe models.Recipe) (int64, error) {
	return s.repo.CreateRecipe(recipe)
}

func (s *RecipeService) GetRecipe(id int64) (models.GetRecipe, error) {
	return s.repo.GetRecipe(id)
}

func (s *RecipeService) GetAllRecipes() ([]models.Recipe, error) {
	return s.repo.GetAllRecipes()
}

func (s *RecipeService) UpdateRecipe(id int64, recipe models.UpdateRecipe) error {
	return s.repo.UpdateRecipe(id, recipe)
}

func (s *RecipeService) DeleteRecipe(id int64) error {
	return s.repo.DeleteRecipe(id)
}
