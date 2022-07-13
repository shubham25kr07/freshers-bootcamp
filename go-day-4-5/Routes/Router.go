package Routes

import (
	"github.com/gin-gonic/gin"
	"go-day-4-5/Controllers"
	"go-day-4-5/Middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	public := r.Group("/api")
	{
		public.POST("/register", Controllers.Register)
		public.POST("/login", Controllers.Login)
	}

	routePath1 := r.Group("/product-api")
	routePath1.Use(Middleware.TokenBasedAuthentication())
	{
		routePath1.GET("products", Controllers.GetProducts)
		routePath1.POST("products", Controllers.AddProduct)
		routePath1.GET("products/:id", Controllers.GetProductById)
		routePath1.PATCH("products/:id", Controllers.UpdateProductDetail)
		routePath1.GET("order/:id", Controllers.GetOrderById)
	}

	routePath2 := r.Group("/customer")
	routePath2.Use(Middleware.TokenBasedAuthentication())
	{
		routePath2.GET("order-history/:customerId", Controllers.GetOrderHistory)
		routePath2.POST("order", Controllers.PlaceOrder)
	}
	return r
}
