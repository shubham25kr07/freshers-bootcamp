package Routes

import (
	"github.com/gin-gonic/gin"
	"q1/Controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	routePath := r.Group("/user-api")
	{
		routePath.GET("user", Controllers.GetUsers)
		routePath.POST("user", Controllers.CreateUser)
		routePath.GET("user/:id", Controllers.GetUserByID)
		routePath.PUT("user/:id", Controllers.UpdateUser)
		routePath.DELETE("user/:id", Controllers.DeleteUser)
	}
	return r
}
