package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-day-4-5/Models"
	"go-day-4-5/Redis"
	"go-day-4-5/Service"
	"net/http"
)

func ModifyProductResponse(product *Models.Product) Models.ProductResponse {
	return Models.ProductResponse{
		ID:          product.ID,
		ProductName: product.ProductName,
		Quantity:    product.Quantity,
		Price:       product.Price,
	}
}

func GetProducts(c *gin.Context) {
	var products []Models.Product
	err := Service.GetAllProducts(&products)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		var productResponse []Models.ProductResponse
		for _, v := range products {
			res := ModifyProductResponse(&v)
			productResponse = append(productResponse, res)
		}
		c.JSON(http.StatusOK, gin.H{"products": productResponse})
	}
}

func AddProduct(c *gin.Context) {
	var product Models.Product
	err2 := c.BindJSON(&product)

	if err2 != nil {
		fmt.Println(err2.Error())
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	err := Service.AddProduct(&product)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		productResponse := ModifyProductResponse(&product)
		c.JSON(http.StatusOK, gin.H{"products": productResponse, "message": "product successfully added"})
	}
}

func GetProductById(c *gin.Context) {
	id := c.Params.ByName("id")
	var product Models.Product
	err := Service.GetProductById(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		productResponse := ModifyProductResponse(&product)
		c.JSON(http.StatusOK, productResponse)
	}
}

func UpdateProductDetail(c *gin.Context) {
	id := c.Params.ByName("id")
	var product Models.Product
	err := Service.GetProductById(&product, id)
	if err != nil {
		c.JSON(http.StatusNotFound, product)
	}
	c.BindJSON(&product)

	uniqueValue := "Locked"
	LockDuration := 6000
	messages := make(chan int, 1)
	Redis.Lock(id, uniqueValue, LockDuration, messages)
	defer Redis.Unlock(id, uniqueValue, messages)

	err = Service.UpdateProduct(&product)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		productResponse := ModifyProductResponse(&product)
		c.JSON(http.StatusOK, productResponse)
	}
	//messages <- true
}
