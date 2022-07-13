package Middleware

import (
	"github.com/gin-gonic/gin"
	"go-day-4-5/Utils"
	"net/http"
)

func TokenBasedAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := Utils.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
