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
		recipe.POST("")
		recipe.GET("")
		recipe.PATCH("")
		recipe.DELETE("")
		ingredient := recipe.Group("/:id/ingredient")
		{
			ingredient.POST("")
			ingredient.GET("")
			ingredient.PATCH("")
			ingredient.DELETE("")
		}
	}

	return router
}
