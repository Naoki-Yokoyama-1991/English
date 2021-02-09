package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/naoki/english/app/service"
	_ "github.com/naoki/english/docs"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// ShowAccount godoc
// @Summary Show a account
// @Description post string by ID
// @Accept  json
// @Produce  json
// @Param  id path int true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} controller.HTTPError
// @Failure 404 {object} controller.HTTPError
// @Failure 500 {object} controller.HTTPError
// @Router /accounts/{id} [get]
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/api")
	v1.POST("/phrases", h.AddPhrase)
	return router
}
