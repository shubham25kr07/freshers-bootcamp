package Controllers

import (
	"github.com/gin-gonic/gin"
	"go-day-4-5/Models"
	"net/http"
)

func Register(c *gin.Context) {

	var input Models.RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := Models.User{}
	user.Username = input.Username
	user.Password = input.Password
	err := Models.SaveUser(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "registration success"})

}

func Login(c *gin.Context) {

	var input Models.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user Models.User
	username := input.Username
	password := input.Password

	token, err := Models.LoginCheck(username, password, &user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}
