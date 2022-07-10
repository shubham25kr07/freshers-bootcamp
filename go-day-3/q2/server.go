package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"q2/Config"
	"q2/Models"
	"q2/Routes"
)

var err error

func main() {

	Config.DB, err = gorm.Open("mysql", Config.DBurl(Config.BuildConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()

	studentModel := &Models.Student{}
	Config.DB.AutoMigrate(studentModel)

	r := Routes.SetupRouter()
	r.Run()
}

// https://www.soberkoder.com/consume-rest-api-go/   good explanation of http put,get, post with golang

// https://golang.cafe/blog/golang-httptest-example.html   good explanation of unit testing of both server and client side

// https://circleci.com/blog/gin-gonic-testing/
