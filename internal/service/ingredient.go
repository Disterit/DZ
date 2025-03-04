package service

import (
	"DZ/internal/repository"
	"DZ/models"
)

type IngredientService struct {
	repo repository.Ingredient
}

func NewIngredientService(repo repository.Ingredient) *IngredientService {
	return &IngredientService{repo: repo}
}

func (s *IngredientService) CreateIngredient(ingredient models.Ingredient) (int64, error) {
	return s.repo.CreateIngredient(ingredient)
}

func (s *IngredientService) UpdateIngredient(id int64, ingredient models.Ingredient) error {
	return s.repo.UpdateIngredient(id, ingredient)
}

func (s *IngredientService) DeleteIngredient(id int64) error {
	return s.repo.DeleteIngredient(id)
}
