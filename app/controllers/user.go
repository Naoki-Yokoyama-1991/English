package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/naoki/gacha/app/models"
	"github.com/naoki/gacha/app/services"
	"github.com/naoki/gacha/app/utils/formaterror"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

// var server = api.Server{}
func (server *Server) CreateUser(c *gin.Context) {

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
	if err != nil {
		errList["Unmarshal_error"] = "Cannot unmarshal body"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	/* @Set services */
	services.Prepare(&user)

	userCreated, err := user.SaveUser(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		errList = formattedError
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":   http.StatusCreated,
			"response": userCreated,
		})
	}
}
