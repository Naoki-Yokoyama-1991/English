package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naoki/gacha/app/api"
	"github.com/naoki/gacha/app/models"
)

var server = api.Server{}

func Like(c *gin.Context) {

	errList := map[string]string{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}

	user := models.User{}

	err = json.Unmarshal(body, &user)
}
