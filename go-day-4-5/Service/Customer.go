package Service

import (
	"go-day-4-5/Config"
	"go-day-4-5/Models"
)

func GetOrderByCustomerId(order *[]Models.Order, customerId string) (err error) {

	if err = Config.DB.Where("customer_id = ?", customerId).Find(order).Error; err != nil {
		return err
	}
	return nil
}

func GetLastOrderForCustomerId(order *Models.Order, customerId string) (err error) {

	if err = Config.DB.Where("customer_id = ?", customerId).Last(order).Error; err != nil {
		return err
	}
	return nil
}
