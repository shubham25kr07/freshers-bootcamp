package Controllers

import (
	"github.com/gin-gonic/gin"
	"go-day-4-5/Models"
	"go-day-4-5/Service"
	"net/http"
)

func GetOrderHistory(c *gin.Context) {
	CustomerId := c.Params.ByName("customerId")
	var order []Models.Order
	err := Service.GetOrderByCustomerId(&order, CustomerId)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		//orderResponse := ModifyResponse(&order)

		c.JSON(http.StatusOK, order)
	}
}
