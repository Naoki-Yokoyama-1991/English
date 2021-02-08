package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/naoki/english/app/service"
	// "github.com/naoki/task/app/controllers"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Routes(r *gin.Engine) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api")
	v1.POST("/phrase", h.AddPhrase)
	return router
}
