package handler

import (
	"DZ/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	recipe := router.Group("/recipe")
	{
		recipe.POST("", h.CreateRecipe)
		recipe.GET("", h.GetAllRecipes)
		recipe.GET("/:id", h.GetRecipe)
		recipe.PATCH("/:id", h.UpdateRecipe)
		recipe.DELETE("/:id", h.DeleteRecipe)
		ingredient := recipe.Group("/:id/ingredient")
		{
			ingredient.POST("", h.CreateIngredient)
		}

		router.PATCH("/ingredient/:id", h.UpdateIngredient)
		router.DELETE("/ingredient/:id", h.DeleteIngredient)
	}

	return router
}
