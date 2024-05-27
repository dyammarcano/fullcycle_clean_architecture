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

func (o *OrderUseCase) UpdateOrder(id int, orderBytes []byte) error {
	order := &domain.Order{}
	if err := json.Unmarshal(orderBytes, order); err != nil {
		return err
	}
	return o.OrderRepo.UpdateOrder(id, order)
}

func (o *OrderUseCase) DeleteOrder(id int) error {
	return o.OrderRepo.DeleteOrder(id)
}

func NewOrderUseCase(repo domain.OrderRepository) *OrderUseCase {
	return &OrderUseCase{OrderRepo: repo}
}
