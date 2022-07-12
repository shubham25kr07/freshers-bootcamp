package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-day-4-5/Models"
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
	err := Models.GetAllProducts(&products)

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

	err := Models.AddProduct(&product)

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
	err := Models.GetProductById(&product, id)
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
	err := Models.GetProductById(&product, id)
	if err != nil {
		c.JSON(http.StatusNotFound, product)
	}
	c.BindJSON(&product)
	err = Models.UpdateProduct(&product)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		productResponse := ModifyProductResponse(&product)
		c.JSON(http.StatusOK, productResponse)
	}
}
