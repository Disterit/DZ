package handler

import (
	"DZ/models"
	"DZ/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateRecipe(c *gin.Context) {

	var input models.Recipe
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logger.Log.Error("Invalid input", err.Error())
		return
	}

	id, err := h.service.Recipe.CreateRecipe(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("Failed to create recipe", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ok": "ok",
		"id": id,
	})
}

func (h *Handler) GetRecipe(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logger.Log.Error("Failed to parse id", err.Error())
		return
	}

	recipe, err := h.service.Recipe.GetRecipe(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("Failed to get recipe", err.Error())
		return
	}

	c.JSON(http.StatusOK, recipe)
}

func (h *Handler) GetAllRecipes(c *gin.Context) {
	recipes, err := h.service.Recipe.GetAllRecipes()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("Failed to get recipes", err.Error())
		return
	}

	c.JSON(http.StatusOK, recipes)
}

func (h *Handler) UpdateRecipe(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logger.Log.Error("Failed to parse id", err.Error())
		return
	}

	var input models.UpdateRecipe
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logger.Log.Error("Invalid input", err.Error())
		return
	}

	err = h.service.Recipe.UpdateRecipe(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("Failed to update recipe", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ok": "ok",
	})
}

func (h *Handler) DeleteRecipe(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logger.Log.Error("Failed to parse id", err.Error())
		return
	}

	err = h.service.Recipe.DeleteRecipe(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("Failed to delete recipe", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ok": "ok",
	})
}
