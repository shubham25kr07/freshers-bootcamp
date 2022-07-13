package Service

import (
	"go-day-4-5/Config"
	"go-day-4-5/Models"
	"go-day-4-5/Redis"
	"strconv"
	"time"
)

func PlaceOrder(order *Models.Order) (err error) {

	var productId = order.ProductId
	uniqueValue := "Locked"
	LockDuration := 6000
	var id = strconv.FormatUint(uint64(productId), 10)
	//messages := make(chan bool)
	if isLocked, _ := Redis.Lock(id, uniqueValue, LockDuration); !isLocked {
		time.Sleep(6000 * time.Millisecond)
	}
	defer Redis.Unlock(id, uniqueValue)

	var product Models.Product
	err1 := GetProductById(&product, id)

	if err1 != nil {
		(*order).Message = "Failed"
	}

	if err1 == nil {
		quantityTobeOrder := order.Quantity
		if quantityTobeOrder > product.Quantity {
			(*order).Message = "Failed"
		} else {

			remainingProductQuantity := product.Quantity - quantityTobeOrder
			// UPDATE PRODUCT QUANTITY
			product.Quantity = remainingProductQuantity
			if err2 := UpdateProduct(&product); err2 != nil {
				(*order).Message = "Failed"
			} else {
				(*order).Message = "order Placed"
			}
		}
	}
	//time.Sleep(6000 * time.Millisecond)
	if err = Config.DB.Create(order).Error; err != nil {
		return err
	}

	return nil
}

func GetOrderById(order *Models.Order, id string) (err error) {

	if err = Config.DB.Where("id = ?", id).First(order).Error; err != nil {
		return err
	}
	return nil
}
