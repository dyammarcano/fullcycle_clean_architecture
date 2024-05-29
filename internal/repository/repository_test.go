package repository

import (
	"github.com/dyammarcano/fullcycle_clean_architecture/internal/entity"
	"testing"
)

func TestRepository(t *testing.T) {
	var repository = Must(NewMemoryRepository())

	var order = &entity.Order{
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

	orderInput := &entity.OrderInputDTO{
		ID:     order.ID,
		Item:   order.Item,
		Amount: order.Amount,
	}

	_, err = repository.UpdateOrder(orderInput.ID, orderInput)
	if err != nil {
		t.Errorf("Error updating order")
	}

	if err = repository.DeleteOrder(orderInput.ID); err != nil {
		t.Errorf("Error deleting order")
	}
}
