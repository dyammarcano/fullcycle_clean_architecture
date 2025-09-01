package usecase

import (
	"github.com/dyammarcano/fullcycle_clean_architecture/internal/domain"
)

type OrderUseCase struct {
	OrderRepo domain.OrderRepository
}

func (o *OrderUseCase) ListOrders() ([]*domain.Order, error) {
	return o.OrderRepo.ListOrders()
}

func (o *OrderUseCase) CreateOrder(order *domain.Order) (*domain.Order, error) {
	return o.OrderRepo.CreateOrder(order)
}

func (o *OrderUseCase) GetOrderByID(id int) (*domain.Order, error) {
	return o.OrderRepo.GetOrderByID(id)
}

func (o *OrderUseCase) UpdateOrder(id int, order *domain.Order) error {
	return o.OrderRepo.UpdateOrder(id, order)
}

func (o *OrderUseCase) DeleteOrder(id int) error {
	return o.OrderRepo.DeleteOrder(id)
}

func NewOrderUseCase(repo domain.OrderRepository) *OrderUseCase {
	return &OrderUseCase{OrderRepo: repo}
}
