package models

import "DZ/pkg/logger"

type Recipe struct {
	Id          int64  `json:"id" db:"id"`
	Name        string `json:"name" db:"name" binding:"required"`
	Author      string `json:"author" db:"author" binding:"required"`
	Description string `json:"description" db:"description"`
}

type GetRecipe struct {
	Recipe      Recipe
	Ingredients []string
}

type UpdateRecipe struct {
	Name        *string `json:"name" db:"name_recipe"`
	Author      *string `json:"author" db:"author"`
	Description *string `json:"description" db:"description"`
}

func (s *UpdateRecipe) Validate() error {
	if s.Name == nil && s.Author == nil && s.Description == nil {
		logger.Log.Error("either title or description must be provided")
		return nil
	}

	return nil
}
