package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"q1/Config"
	"q1/Model"
	"q1/Routes"
)

var err error

func main() {

	Config.DB, err = gorm.Open("mysql", Config.DBurl(Config.BuildConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()

	userModel := &Model.User{}
	Config.DB.AutoMigrate(userModel)

	r := Routes.SetupRouter()
	r.Run()
}
