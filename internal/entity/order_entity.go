package entity

import "errors"

// ErrInvalidEntity is used when an entity is invalid
var ErrInvalidEntity = errors.New("invalid entity")

var _ CreateOrderUseCase = (*createOrderUseCase)(nil)

type CreateOrderUseCase interface {
	Execute(input *OrderInputDTO) (*OrderOutputDTO, error)
}

type createOrderUseCase struct {
	repo OrderRepository
}

func (c *createOrderUseCase) Execute(input *OrderInputDTO) (*OrderOutputDTO, error) {
	order := Order{
		ID:     input.ID,
		Item:   input.Item,
		Amount: input.Amount,
	}
	if err := order.IsValid(); err != nil {
		return nil, err
	}
	createdOrder, err := c.repo.CreateOrder(&order)
	if err != nil {
		return nil, err
	}
	return createdOrder, nil
}

func NewCreateOrderUseCase(repo OrderRepository) CreateOrderUseCase {
	return &createOrderUseCase{
		repo: repo,
	}
}

type OrderRepository interface {
	ListOrders() ([]*OrderOutputDTO, error)
	CreateOrder(order *Order) (*OrderOutputDTO, error)
	GetOrderByID(id int) (*OrderOutputDTO, error)
	UpdateOrder(id int, order *OrderInputDTO) (*OrderOutputDTO, error)
	DeleteOrder(id int) error
}

type OrderInputDTO struct {
	ID     int     `json:"id"`
	Item   string  `json:"item"`
	Amount float32 `json:"amount"`
}

type OrderOutputDTO struct {
	ID     int     `json:"id"`
	Item   string  `json:"item"`
	Amount float32 `json:"amount"`
}

type Order struct {
	ID     int
	Item   string
	Amount float32
}

func NewOrder(id int, item string, amount float32) (*Order, error) {
	order := &Order{
		ID:     id,
		Item:   item,
		Amount: amount,
	}
	if err := order.IsValid(); err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) IsValid() error {
	if o.Item == "" {
		return ErrInvalidEntity
	}
	if o.Amount <= 0 {
		return ErrInvalidEntity
	}
	return nil
}
