package service

import "DZ/internal/repository"

type Ingredient interface {
}

type Recipe interface {
}

type Service struct {
	Recipe     Recipe
	Ingredient Ingredient
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
