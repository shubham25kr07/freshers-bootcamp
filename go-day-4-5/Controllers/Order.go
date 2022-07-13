package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-day-4-5/Models"
	"go-day-4-5/Service"
	"net/http"
	"strconv"
	"time"
)

func ModifyResponse(order *Models.Order) Models.OrderResponse {
	return Models.OrderResponse{
		ID:        order.ID,
		ProductId: order.ProductId,
		Quantity:  order.Quantity,
		Message:   order.Message,
	}
}

func PlaceOrder(c *gin.Context) {

	var order Models.Order
	c.BindJSON(&order)

	customerId := order.CustomerId
	var lastOrder Models.Order
	cid := strconv.FormatUint(uint64(customerId), 10)
	err1 := Service.GetLastOrderForCustomerId(&lastOrder, cid)

	lastTime, _ := strconv.Atoi(strconv.FormatInt(lastOrder.CreatedAt.UTC().Unix(), 10))
	presentTime, _ := strconv.Atoi(strconv.FormatInt(time.Now().UTC().Unix(), 10))
	if presentTime-lastTime < 30 && err1 == nil {
		c.JSON(http.StatusOK, gin.H{"message": "transaction in cool down period of 5 min"})
		return
	}
	err := Service.PlaceOrder(&order)
	orderResponse := ModifyResponse(&order)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, orderResponse)
	}
}

func GetOrderById(c *gin.Context) {
	id := c.Params.ByName("id")
	var order Models.Order
	err := Service.GetOrderById(&order, id)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		orderResponse := ModifyResponse(&order)
		c.JSON(http.StatusOK, orderResponse)
	}
}
