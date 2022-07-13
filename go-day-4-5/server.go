package main

import (
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-day-4-5/Config"
	"go-day-4-5/Models"
	"go-day-4-5/Routes"
)

var err error

func main() {

	Config.DB, err = gorm.Open("mysql", Config.DBurl(Config.BuildConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()

	flag.Parse()
	Config.Pool = Config.NewPool(*Config.RedisServer)

	orderModel := &Models.Order{}
	productModel := &Models.Product{}
	customerModel := &Models.Customer{}
	userModel := &Models.User{}
	Config.DB.AutoMigrate(orderModel, productModel, customerModel, userModel)

	r := Routes.SetupRouter()
	r.Run()
}
