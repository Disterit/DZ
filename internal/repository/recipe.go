package repository

import (
	"DZ/models"
	"DZ/pkg/logger"
	"database/sql"
	"fmt"
	"strings"
)

type RecipeRepository struct {
	db *sql.DB
}

func NewRecipeRepository(db *sql.DB) *RecipeRepository {
	return &RecipeRepository{db: db}
}

func (r *RecipeRepository) CreateRecipe(recipe models.Recipe) (int64, error) {

	var id int64

	row := r.db.QueryRow("INSERT INTO recipe (name_recipe, author, description) VALUES ($1, $2, $3) RETURNING id", recipe.Name, recipe.Author, recipe.Description)
	if err := row.Scan(&id); err != nil {
		logger.Log.Error("Error inserting recipe into database", err)
		return 0, err
	}

	return id, nil
}

func (r *RecipeRepository) GetRecipe(id int64) (models.GetRecipe, error) {
	var recipe models.GetRecipe

	row, err := r.db.Query("SELECT * FROM recipe WHERE id=$1", id)
	if err != nil {
		logger.Log.Error("Error selecting recipe from database", err)
		return models.GetRecipe{}, err
	}

	for row.Next() {
		err = row.Scan(&recipe.Recipe.Id, &recipe.Recipe.Name, &recipe.Recipe.Author, &recipe.Recipe.Description)
		if err != nil {
			logger.Log.Error("Error selecting recipe from database", err)
			return models.GetRecipe{}, err
		}
	}

	row, err = r.db.Query("SELECT (ingredient.name_ingredient) FROM recipe JOIN ingredient ON recipe.id = ingredient.id_recipe WHERE recipe.id = $1", id)
	if err != nil {
		logger.Log.Error("Error selecting ingredient from database", err)
		return models.GetRecipe{}, err
	}

	for row.Next() {
		var ingredient string
		err = row.Scan(&ingredient)
		if err != nil {
			logger.Log.Error("Error selecting ingredient from database", err)
			return models.GetRecipe{}, err
		}
		recipe.Ingredients = append(recipe.Ingredients, ingredient)
	}

	return recipe, nil
}

func (r *RecipeRepository) GetAllRecipes() ([]models.Recipe, error) {
	var recipes []models.Recipe

	rows, err := r.db.Query("SELECT * FROM recipe")
	if err != nil {
		logger.Log.Error("Error selecting recipes from database", err)
		return []models.Recipe{}, err
	}

	for rows.Next() {
		var recipe models.Recipe
		if err = rows.Scan(&recipe.Id, &recipe.Name, &recipe.Author, &recipe.Description); err != nil {
			logger.Log.Error("Error selecting recipes from database", err)
			return []models.Recipe{}, err
		}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func (r *RecipeRepository) UpdateRecipe(id int64, recipe models.UpdateRecipe) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if recipe.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name_recipe=$%d", argId))
		args = append(args, *recipe.Name)
		argId++
	}

	if recipe.Author != nil {
		setValues = append(setValues, fmt.Sprintf("author=$%d", argId))
		args = append(args, *recipe.Author)
		argId++
	}

	if recipe.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *recipe.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ",")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", recipeTable, setQuery, argId)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		logger.Log.Error("Error updating recipe from database", err)
		return err
	}

	return nil
}

func (r *RecipeRepository) DeleteRecipe(id int64) error {
	_, err := r.db.Exec("DELETE FROM recipe WHERE id=$1", id)
	if err != nil {
		logger.Log.Error("Error deleting recipe from database", err)
		return err
	}

	return nil
}
