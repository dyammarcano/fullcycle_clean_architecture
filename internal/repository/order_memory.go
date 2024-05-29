package repository

import (
	"github.com/dyammarcano/fullcycle_clean_architecture/internal/entity"
)

type OrderMemoryRepository struct {
	orders map[int]*entity.OrderOutputDTO
}

func (r *OrderMemoryRepository) GetOrderByID(id int) (*entity.OrderOutputDTO, error) {
	return r.orders[id], nil
}

func (r *OrderMemoryRepository) ListOrders() ([]*entity.OrderOutputDTO, error) {
	orders := make([]*entity.OrderOutputDTO, 0)
	for _, order := range r.orders {
		orders = append(orders, order)
	}
	return orders, nil
}

func (r *OrderMemoryRepository) CreateOrder(order *entity.Order) (*entity.OrderOutputDTO, error) {
	if order == nil {
		return nil, ErrInvalidEntity
	}
	if _, ok := r.orders[order.ID]; ok {
		return nil, ErrUserAlreadyExists
	}
	order.ID = len(r.orders) + 1
	r.orders[order.ID] = &entity.OrderOutputDTO{
		ID:     order.ID,
		Item:   order.Item,
		Amount: order.Amount,
	}
	return r.orders[order.ID], nil
}

func (r *OrderMemoryRepository) UpdateOrder(id int, order *entity.OrderInputDTO) (*entity.OrderOutputDTO, error) {
	r.orders[id] = &entity.OrderOutputDTO{
		ID:     order.ID,
		Item:   order.Item,
		Amount: order.Amount,
	}
	return r.orders[order.ID], nil
}

func (r *OrderMemoryRepository) DeleteOrder(id int) error {
	delete(r.orders, id)
	return nil
}

func (r *OrderMemoryRepository) Close() error {
	r.orders = nil
	return nil
}

func NewMemoryRepository() (entity.OrderRepository, error) {
	return &OrderMemoryRepository{orders: make(map[int]*entity.OrderOutputDTO)}, nil
}
