package server

import (
	"github.com/bala-golang/GoWebhookTransformer/api/handlers"
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Recovery())

	// Routes
	router.POST("/process-event", handlers.ProcessEventHandler)

	return router
}
