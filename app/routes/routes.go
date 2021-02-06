package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/naoki/task/app/controllers"
)

func Routes(r *gin.Engine) {

	v1 := r.Group("/api")
	v1.POST("/users", controllers.CreateUser)
}
