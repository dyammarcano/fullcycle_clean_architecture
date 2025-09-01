package repository

import (
	"testing"

	"github.com/dyammarcano/fullcycle_clean_architecture/internal/domain"
)

func TestRepository(t *testing.T) {
	repository, err := NewMemoryRepository()
	if err != nil {
		t.Fatalf("Error creating repository")
	}

	var order = &domain.Order{
		Item:   "Bag",
		Amount: 2,
	}

	if _, err := repository.CreateOrder(order); err != nil {
		t.Errorf("Error creating order")
	}

	orders, err := repository.ListOrders()
	if err != nil {
		t.Errorf("Error listing orders")
	}

	if len(orders) == 0 {
		t.Errorf("Error listing orders")
	}

	if _, err = repository.GetOrderByID(order.ID); err != nil {
		t.Errorf("Error getting order")
	}

	order.Amount = 5

	if err = repository.UpdateOrder(order.ID, order); err != nil {
		t.Errorf("Error updating order")
	}

	if err = repository.DeleteOrder(order.ID); err != nil {
		t.Errorf("Error deleting order")
	}
}
