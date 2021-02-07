package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naoki/english/app/models"
	"github.com/naoki/english/app/service"
)

type Handler struct {
	services *service.Service
}

func (h *Handler) AddPhrase(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Unable to get request")
		return
	}

	var phrase models.Phrase
	err = json.Unmarshal(body, &phrase)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Cannot unmarshal body")
		return
	}

}
