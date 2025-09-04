package usecase

import (
	"encoding/json"

	"github.com/dyammarcano/fullcycle_clean_architecture/internal/domain"
)

type OrderUseCase struct {
	OrderRepo domain.OrderRepository
}

func (o *OrderUseCase) ListOrders() ([]*domain.Order, error) {
	return o.OrderRepo.ListOrders()
}

func (o *OrderUseCase) CreateOrder(orderBytes []byte) (*domain.Order, error) {
	order := &domain.Order{}
	if err := json.Unmarshal(orderBytes, order); err != nil {
		return nil, err
	}

	return o.OrderRepo.CreateOrder(order)
}

func (o *OrderUseCase) GetOrderByID(id int) (*domain.Order, error) {
	return o.OrderRepo.GetOrderByID(id)
}

func (o *OrderUseCase) UpdateOrder(id int, orderBytes []byte) (*domain.Order, error) {
	order := &domain.Order{}
	if err := json.Unmarshal(orderBytes, order); err != nil {
		return nil, err
	}

	if err := o.OrderRepo.UpdateOrder(id, order); err != nil {
		return nil, err
	}

	order.ID = id

	return order, nil
}

func (o *OrderUseCase) DeleteOrder(id int) error {
	return o.OrderRepo.DeleteOrder(id)
}

func NewOrderUseCase(repo domain.OrderRepository) *OrderUseCase {
	return &OrderUseCase{OrderRepo: repo}
}
