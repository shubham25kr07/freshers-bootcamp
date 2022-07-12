package Models

import "time"

type Order struct {
	ID         uint32 `json:"order_id"`
	CustomerId uint32 `json:"customer_id"`
	ProductId  uint32 `json:"product_id"`
	Quantity   uint16 `json:"quantity"`
	Message    string `json:"message"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Product    Product
	Customer   Customer
}

func (o *Order) TableName() string {
	return "orders"
}

type OrderResponse struct {
	ID        uint32 `json:"order_id"`
	ProductId uint32 `json:"product_id"`
	Quantity  uint16 `json:"quantity"`
	Message   string `json:"message"`
}
