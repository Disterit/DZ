package handler

import (
	"DZ/models"
	"DZ/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateIngredient(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid RecipeID")
		logger.Log.Error("CreateIngredient Error: ", err.Error())
		return
	}

	var input models.Ingredient
	input.RecipeId = id
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid Recipe")
		logger.Log.Error("CreateIngredient Error: ", err.Error())
		return
	}

	idIng, err := h.service.Ingredient.CreateIngredient(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("CreateIngredient Error: ", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ok": "ok",
		"id": idIng,
	})
}

func (h *Handler) UpdateIngredient(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid RecipeID")
		logger.Log.Error("UpdateIngredient Error: ", err.Error())
		return
	}

	var input models.Ingredient
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid Recipe")
		logger.Log.Error("UpdateIngredient Error: ", err.Error())
		return
	}

	err = h.service.Ingredient.UpdateIngredient(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("UpdateIngredient Error: ", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ok": "ok",
	})
}

func (h *Handler) DeleteIngredient(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid RecipeID")
		logger.Log.Error("DeleteIngredient Error: ", err.Error())
		return
	}

	err = h.service.Ingredient.DeleteIngredient(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("DeleteIngredient Error: ", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ok": "ok",
	})
}
