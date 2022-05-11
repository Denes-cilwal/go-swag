package route

import (
	"go-swag/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	user := r.Group("/user")
	{
		user.GET("/", controller.GetUsers)
		user.POST("/", controller.CreateUser)
		user.GET("/:id", controller.GetUserByID)
	}
	return r
}
