package domain

import (
	"encoding/json"
)

type OrderRepository interface {
	ListOrders() ([]*Order, error)
	CreateOrder(order *Order) (*Order, error)
	GetOrderByID(id int) (*Order, error)
	UpdateOrder(id int, order *Order) error
	DeleteOrder(id int) error
}

type Order struct {
	ID     int     `json:"id"`
	Item   string  `json:"item"`
	Amount float32 `json:"amount"`
}

func (o *Order) Bytes() []byte {
	data, err := json.Marshal(o)
	if err != nil {
		return nil
	}
	return data
}
