package Models

import "time"

type Customer struct {
	ID        uint32 `json:"customer_id"`
	Name      string `json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *Customer) TableName() string {
	return "customers"
}
