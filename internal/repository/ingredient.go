package repository

import (
	"DZ/models"
	"DZ/pkg/logger"
	"database/sql"
)

type IngredientRepository struct {
	db *sql.DB
}

func NewIngredientRepository(db *sql.DB) *IngredientRepository {
	return &IngredientRepository{db: db}
}

func (i *IngredientRepository) CreateIngredient(ingredient models.Ingredient) (int64, error) {

	var id int64

	row := i.db.QueryRow("INSERT INTO ingredient (id_recipe, name_ingredient) VALUES ($1, $2);", ingredient.RecipeId, ingredient.Name)
	if err := row.Scan(&id); err != nil {
		logger.Log.Error("error scanning ingredient row:", err)
		return 0, err
	}

	return id, nil
}

func (i *IngredientRepository) UpdateIngredient(id int64, ingredient models.Ingredient) error {

	_, err := i.db.Exec("UPDATE ingredient SET name_ingredient = $1 WHERE id = $2;", ingredient.Name, id)
	if err != nil {
		logger.Log.Error("error updating ingredient row:", err)
		return err
	}

	return nil
}

func (i *IngredientRepository) DeleteIngredient(id int64) error {

	_, err := i.db.Exec("DELETE FROM ingredient WHERE id = $1;", id)
	if err != nil {
		logger.Log.Error("error deleting ingredient row:", err)
		return err
	}

	return nil
}
