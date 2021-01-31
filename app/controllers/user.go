package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/naoki/gacha/app/api"
)

var server = api.Server{}

func CreateUser(c *gin.Context) {

	/* this is here
	errList := map[string]string{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
	}
	*/
}
