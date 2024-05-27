package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/dyammarcano/fullcycle_clean_architecture/internal/domain"
	"github.com/dyammarcano/fullcycle_clean_architecture/pkg/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"time"
)

type OrderSqliteRepository struct {
	db *sql.DB
}

func (r *OrderSqliteRepository) ListOrders() ([]*domain.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, "SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make([]*domain.Order, 0)
	for rows.Next() {
		var order domain.Order
		if err = rows.Scan(&order.ID, &order.Item, &order.Amount); err != nil {
			return orders, err
		}
		orders = append(orders, &order)
	}

	return orders, nil
}

func (r *OrderSqliteRepository) CreateOrder(order *domain.Order) (*domain.Order, error) {
	if order == nil {
		return nil, ErrInvalidEntity
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err := r.db.PrepareContext(ctx, `INSERT INTO orders(item, amount) VALUES($1, $2)`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(order.Item, order.Amount)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	order.ID = int(id)

	return order, nil
}

func (r *OrderSqliteRepository) GetOrderByID(id int) (*domain.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := r.db.QueryRowContext(ctx, "SELECT * FROM orders WHERE id = ?", id)

	var data string
	if err := row.Scan(&data); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrOrderNotFound
		}
		return nil, err
	}

	order := &domain.Order{}
	if err := json.Unmarshal([]byte(data), order); err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderSqliteRepository) UpdateOrder(id int, order *domain.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err := r.db.PrepareContext(ctx, `UPDATE orders SET item = $1, amount = $2, WHERE id = $3`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(order.Item, order.Amount, id); err != nil {
		return err
	}

	return nil
}

func (r *OrderSqliteRepository) DeleteOrder(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err := r.db.PrepareContext(ctx, "DELETE FROM orders WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrOrderNotFound
	}

	return nil
}

func NewOrderSqliteRepository() (domain.OrderRepository, error) {
	db, err := sql.Open("sqlite3", config.G.Db.DBPath)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(config.G.Db.MaxOpenConns)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance("file://internal/repository/migrations", "", driver)
	if err != nil {
		return nil, err
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, err
	}

	return &OrderSqliteRepository{db: db}, nil
}
