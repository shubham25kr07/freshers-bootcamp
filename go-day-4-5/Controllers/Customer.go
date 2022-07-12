package Controllers

import (
	"github.com/gin-gonic/gin"
	"go-day-4-5/Models"
	"net/http"
)

func GetOrderHistory(c *gin.Context) {
	CustomerId := c.Params.ByName("customerId")
	var order []Models.Order
	err := Models.GetOrderByCustomerId(&order, CustomerId)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		//orderResponse := ModifyResponse(&order)

		c.JSON(http.StatusOK, order)
	}
}
