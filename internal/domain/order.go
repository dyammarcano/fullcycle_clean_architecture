package domain

type OrderRepository interface {
	ListOrders() ([]*Order, error)
	CreateOrder(order *Order) (*Order, error)
	GetOrderByID(id int) (*Order, error)
	UpdateOrder(id int, order *Order) error
	DeleteOrder(id int) error
}

type Order struct {
	ID     int
	Item   string
	Amount float32
}
