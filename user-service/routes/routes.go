package routes

import (
	"user-service/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/users/:id", controllers.GetUserByID)
	return r
}
