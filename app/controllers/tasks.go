package controllers

import (
	"github.com/gin-gonic/gin"
)

// var server = api.Server{}
func CreateTask(c *gin.Context) {

	// errList := map[string]string{}
	// body, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	errList["Invalid_body"] = "Unable to get request"
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"status": http.StatusBadRequest,
	// 		"error":  errList,
	// 	})
	// 	return
	// }

	// user := models.User{}
	// err = json.Unmarshal(body, &user)
	// if err != nil {
	// 	errList["Unmarshal_error"] = "Cannot unmarshal body"
	// 	c.JSON(http.StatusUnprocessableEntity, gin.H{
	// 		"status": http.StatusUnprocessableEntity,
	// 		"error":  errList,
	// 	})
	// 	return
	// }

	// /* @Set services */
	// services.Prepare(&user)
	// errorMessages := validate.Validate(&user, "")
	// if len(errorMessages) > 0 {
	// 	errList = errorMessages
	// 	c.JSON(http.StatusUnprocessableEntity, gin.H{
	// 		"status": http.StatusUnprocessableEntity,
	// 		"error":  errList,
	// 	})
	// 	return
	// }

	// userCreated, err := user.SaveUser(db.DB)
	// if err != nil {
	// 	formattedError := formaterror.FormatError(err.Error())
	// 	errList = formattedError
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"status":   http.StatusCreated,
	// 		"response": userCreated,
	// 	})
	// } else {
	// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "user": user})
	// }
}

// func GetUser(c *gin.Context) {
// 	errList = map[string]string{}

// 	userID := c.Param("id")

// 	uid, err := strconv.ParseUint(userID, 10, 32)
// 	if err != nil {

// 	}
// }
