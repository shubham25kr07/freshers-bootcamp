package Routes

import (
	"github.com/gin-gonic/gin"
	"go-day-4-5/Controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	routePath1 := r.Group("/product-api")
	{
		routePath1.GET("products", Controllers.GetProducts)
		routePath1.POST("products", Controllers.AddProduct)
		routePath1.GET("products/:id", Controllers.GetProductById)
		routePath1.PATCH("products/:id", Controllers.UpdateProductDetail)
	}
	routePath2 := r.Group("/order-api")
	{
		routePath2.GET("order/:id", Controllers.GetOrderById)
	}
	routePath3 := r.Group("/customer")
	{
		routePath3.GET("order-history/:customerId", Controllers.GetOrderHistory)
		routePath3.POST("order", Controllers.PlaceOrder)
	}
	return r
}
