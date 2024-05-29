package usecase

import (
	"encoding/json"
	"github.com/dyammarcano/fullcycle_clean_architecture/internal/entity"
)

type OrderUseCase struct {
	OrderRepo entity.OrderRepository
}

func (o *OrderUseCase) ListOrders() ([]*entity.OrderOutputDTO, error) {
	return o.OrderRepo.ListOrders()
}

func (o *OrderUseCase) CreateOrder(orderBytes []byte) (*entity.OrderOutputDTO, error) {
	order := &entity.Order{}
	if err := json.Unmarshal(orderBytes, order); err != nil {
		return nil, err
	}
	if err := order.IsValid(); err != nil {
		return nil, err
	}
	return o.OrderRepo.CreateOrder(order)
}

func (o *OrderUseCase) GetOrderByID(id int) (*entity.OrderOutputDTO, error) {
	return o.OrderRepo.GetOrderByID(id)
}

func (o *OrderUseCase) UpdateOrder(id int, orderBytes []byte) (*entity.OrderOutputDTO, error) {
	order := &entity.OrderInputDTO{}
	if err := json.Unmarshal(orderBytes, order); err != nil {
		return nil, err
	}
	return o.OrderRepo.UpdateOrder(id, order)
}

func (o *OrderUseCase) DeleteOrder(id int) error {
	return o.OrderRepo.DeleteOrder(id)
}

func NewOrderUseCase(repo entity.OrderRepository) *OrderUseCase {
	return &OrderUseCase{OrderRepo: repo}
}
