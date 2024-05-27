package repository

import (
	"github.com/dyammarcano/fullcycle_clean_architecture/internal/domain"
)

type OrderMemoryRepository struct {
	orders map[int]*domain.Order
}

func (r *OrderMemoryRepository) GetOrderByID(id int) (*domain.Order, error) {
	return r.orders[id], nil
}

func (r *OrderMemoryRepository) ListOrders() ([]*domain.Order, error) {
	orders := make([]*domain.Order, 0)
	for _, order := range r.orders {
		orders = append(orders, order)
	}
	return orders, nil
}

func (r *OrderMemoryRepository) CreateOrder(order *domain.Order) (*domain.Order, error) {
	if order == nil {
		return nil, ErrInvalidEntity
	}
	if _, ok := r.orders[order.ID]; ok {
		return nil, ErrUserAlreadyExists
	}
	order.ID = len(r.orders) + 1
	r.orders[order.ID] = order
	return order, nil
}

func (r *OrderMemoryRepository) UpdateOrder(id int, order *domain.Order) error {
	r.orders[id] = order
	return nil
}

func (r *OrderMemoryRepository) DeleteOrder(id int) error {
	delete(r.orders, id)
	return nil
}

func (r *OrderMemoryRepository) Close() error {
	r.orders = nil
	return nil
}

func NewMemoryRepository() (domain.OrderRepository, error) {
	return &OrderMemoryRepository{orders: make(map[int]*domain.Order)}, nil
}
