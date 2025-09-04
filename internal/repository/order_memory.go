package repository

import (
	"errors"

	"github.com/dyammarcano/fullcycle_clean_architecture/internal/domain"
)

type OrderMemoryRepository struct {
	orders map[int]*domain.Order
}

func (r *OrderMemoryRepository) GetOrderByID(id int) (*domain.Order, error) {
	order, ok := r.orders[id]
	if !ok {
		return nil, errors.New("order not found")
	}

	return order, nil
}

func (r *OrderMemoryRepository) ListOrders() ([]*domain.Order, error) {
	orders := make([]*domain.Order, 0, len(r.orders))
	for _, order := range r.orders {
		orders = append(orders, order)
	}

	return orders, nil
}

func (r *OrderMemoryRepository) CreateOrder(order *domain.Order) (*domain.Order, error) {
	if order == nil {
		return nil, errors.New("invalid entity")
	}

	if order.ID != 0 {
		if _, ok := r.orders[order.ID]; ok {
			return nil, errors.New("order already exists")
		}
	}

	order.ID = len(r.orders) + 1
	r.orders[order.ID] = order

	return order, nil
}

func (r *OrderMemoryRepository) UpdateOrder(id int, order *domain.Order) error {
	if order == nil {
		return errors.New("invalid entity")
	}

	if _, ok := r.orders[id]; !ok {
		return errors.New("order not found")
	}

	order.ID = id
	r.orders[id] = order

	return nil
}

func (r *OrderMemoryRepository) DeleteOrder(id int) error {
	if _, ok := r.orders[id]; !ok {
		return errors.New("order not found")
	}

	delete(r.orders, id)

	return nil
}

func (r *OrderMemoryRepository) Close() error {
	r.orders = nil

	return nil
}

func NewOrderMemoryRepository() (domain.OrderRepository, error) {
	return &OrderMemoryRepository{
		orders: make(map[int]*domain.Order),
	}, nil
}

// NewMemoryRepository kept for backward compatibility in tests
func NewMemoryRepository() (domain.OrderRepository, error) {
	return NewOrderMemoryRepository()
}
