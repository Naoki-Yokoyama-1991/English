package routes

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	v1 := r.Group("/api")
	v1.GET("/")
	v1.POST("/")
}
