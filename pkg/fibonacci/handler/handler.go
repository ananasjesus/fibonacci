package handler

import (
	"github.com/ananasjesus/fibonacci/pkg/cache"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	cache cache.Cache
}

func NewHandler(cache cache.Cache) *Handler {
	return &Handler{cache: cache}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.GET("/fibonacci", h.fibonacci)
	}

	return router
}
