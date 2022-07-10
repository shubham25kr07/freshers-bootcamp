package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"q2/Config"
	"q2/Controllers"
	"q2/Models"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}
func TestGetUsers(t *testing.T) {

	Config.DB, err = gorm.Open("mysql", Config.DBurl(Config.BuildConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	studentModel := &Models.Student{}
	Config.DB.AutoMigrate(studentModel)

	r := SetUpRouter()
	routePath := r.Group("/student-api")
	routePath.GET("student", Controllers.GetUsers)

	req, _ := http.NewRequest("GET", "/student-api/student", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	//responseData, _ := ioutil.ReadAll(w.Body)
	//sb := string(responseData)
	//log.Printf(sb)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetUserByID(t *testing.T) {

	Config.DB, err = gorm.Open("mysql", Config.DBurl(Config.BuildConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	studentModel := &Models.Student{}
	Config.DB.AutoMigrate(studentModel)

	r := SetUpRouter()
	routePath := r.Group("/student-api")
	routePath.GET("student/:id", Controllers.GetUserByID)

	dataId := []struct {
		id string
	}{
		{"1"},
		{"7"},
		{"2"},
	}
	for _, Id := range dataId {
		req, _ := http.NewRequest("GET", "/student-api/student/"+Id.id, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		//responseData, _ := ioutil.ReadAll(w.Body)
		//sb := string(responseData)
		//log.Printf(sb)
		assert.Equal(t, http.StatusOK, w.Code)
	}
}
