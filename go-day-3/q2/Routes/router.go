package Routes

import (
	"github.com/gin-gonic/gin"
	"q2/Controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	routePath := r.Group("/student-api")
	{
		routePath.GET("student", Controllers.GetUsers)
		routePath.POST("student", Controllers.CreateUser)
		routePath.GET("student/:id", Controllers.GetUserByID)
		routePath.PUT("student/:id", Controllers.UpdateUser)
		//routePath.DELETE("student/:id", Controllers.DeleteUser)
	}
	return r
}
