package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naoki/english/app/models"
)

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

	phraseCreate, err := h.services.Phrase.Create(&phrase)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":   http.StatusCreated,
		"response": phraseCreate,
	})

}
