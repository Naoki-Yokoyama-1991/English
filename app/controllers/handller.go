package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/naoki/english/app/service"
)

type Phrase interface {
	AddPhrase(c *gin.Context)
}
type Handler struct {
	Phrase
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		Phrase: NewPhraseHandler(services.Phrase),
	}
}

func (h *Handler) Handler() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("/api")
	v1.POST("/phrases", h.AddPhrase)
	return router
}
