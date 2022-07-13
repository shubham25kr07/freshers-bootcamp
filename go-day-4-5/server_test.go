package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"go-day-4-5/Config"
	"go-day-4-5/Controllers"
	"go-day-4-5/Models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}
func TestGetProducts(t *testing.T) {

	Config.DB, err = gorm.Open("mysql", Config.DBurl(Config.BuildConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()

	productModel := &Models.Product{}
	Config.DB.AutoMigrate(productModel)

	flag.Parse()
	Config.Pool = Config.NewPool(*Config.RedisServer)
	
	r := SetUpRouter()
	routePath := r.Group("/product-api")
	routePath.GET("products", Controllers.GetProducts)

	req, _ := http.NewRequest("GET", "/product-api/products", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	//responseData, _ := ioutil.ReadAll(w.Body)
	//sb := string(responseData)
	//log.Printf(sb)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetProductById(t *testing.T) {

	Config.DB, err = gorm.Open("mysql", Config.DBurl(Config.BuildConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	productModel := &Models.Product{}
	Config.DB.AutoMigrate(productModel)

	r := SetUpRouter()
	routePath := r.Group("/product-api")
	routePath.GET("products/:id", Controllers.GetProductById)

	dataId := []struct {
		id string
	}{
		{"1"},
		{"7"},
		{"2"},
	}
	for _, Id := range dataId {
		req, _ := http.NewRequest("GET", "/product-api/products/"+Id.id, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		//responseData, _ := ioutil.ReadAll(w.Body)
		//sb := string(responseData)
		//log.Printf(sb)
		assert.Equal(t, http.StatusOK, w.Code)
	}
}

func TestGetOrderById(t *testing.T) {

	Config.DB, err = gorm.Open("mysql", Config.DBurl(Config.BuildConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	orderModel := &Models.Order{}
	Config.DB.AutoMigrate(orderModel)

	r := SetUpRouter()
	routePath := r.Group("/product-api")
	routePath.GET("order/:id", Controllers.GetProductById)

	dataId := []struct {
		id string
	}{
		{"1"},
		{"7"},
		{"2"},
	}
	for _, Id := range dataId {
		req, _ := http.NewRequest("GET", "/product-api/order/"+Id.id, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		//responseData, _ := ioutil.ReadAll(w.Body)
		//sb := string(responseData)
		//log.Printf(sb)
		assert.Equal(t, http.StatusOK, w.Code)
	}
}
