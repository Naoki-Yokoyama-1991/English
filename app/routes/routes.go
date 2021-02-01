package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/naoki/gacha/app/controllers"
)

func Routes(r *gin.Engine) {

	v2 := r.Group("/api")
	v2.POST("/users", controllers.CreateUser)
}
