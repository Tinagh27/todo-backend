package http

import (
	"todo-backend/internal/adapter/inbound/http/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	healthHandler := handler.NewHealthHandler()

	router.GET("/health", healthHandler.Check)

	return router
}
