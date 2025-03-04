package models

type Recipe struct {
	Id          int64  `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Author      string `json:"author" db:"author"`
	Description string `json:"description" db:"description"`
}

type Ingredient struct {
	Id       int64  `json:"id" db:"id"`
	RecipeId int64  `json:"recipe_id" db:"recipe_id"`
	Name     string `json:"name" db:"name"`
}
