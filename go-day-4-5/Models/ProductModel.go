package Models

import "time"

type Product struct {
	ID          uint32 `json:"product_id"`
	ProductName string `json:"product_name"`
	Price       uint16 `json:"price"`
	Quantity    uint16 `json:"quantity"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (p *Product) TableName() string {
	return "products"
}

type ProductResponse struct {
	ID          uint32 `json:"product_id"`
	ProductName string `json:"product_name"`
	Price       uint16 `json:"price"`
	Quantity    uint16 `json:"quantity"`
}
