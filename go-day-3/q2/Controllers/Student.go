package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"q2/Models"
)

func GetUsers(c *gin.Context) {
	var student []Models.Student
	err := Models.GetAllUsers(&student)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"data": student})
	}
}

func CreateUser(c *gin.Context) {
	var student Models.Student
	c.BindJSON(&student)
	err := Models.CreateUser(&student)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var student Models.Student
	err := Models.GetUserByID(&student, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}
func UpdateUser(c *gin.Context) {
	var student Models.Student
	id := c.Params.ByName("id")
	err := Models.GetUserByID(&student, id)
	if err != nil {
		c.JSON(http.StatusNotFound, student)
	}
	c.BindJSON(&student)
	err = Models.UpdateUser(&student)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}
