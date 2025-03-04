package models

type Ingredient struct {
	Id       int64  `json:"id" db:"id"`
	RecipeId int64  `json:"recipe_id" db:"recipe_id" binding:"required"`
	Name     string `json:"name" db:"name" binding:"required"`
}
