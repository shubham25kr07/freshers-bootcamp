package Models

import "go-day-4-5/Config"

func GetOrderByCustomerId(order *[]Order, customerId string) (err error) {

	if err = Config.DB.Where("customer_id = ?", customerId).Find(order).Error; err != nil {
		return err
	}
	return nil
}

func GetLastOrderForCustomerId(order *Order, customerId string) (err error) {

	if err = Config.DB.Where("customer_id = ?", customerId).Last(order).Error; err != nil {
		return err
	}
	return nil
}
