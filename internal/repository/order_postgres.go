package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dyammarcano/fullcycle_clean_architecture/internal/entity"
	"github.com/dyammarcano/fullcycle_clean_architecture/pkg/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"time"
)

type OrderPostgresRepository struct {
	db *sql.DB
}

func (r *OrderPostgresRepository) ListOrders() ([]*entity.OrderOutputDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, "SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make([]*entity.OrderOutputDTO, 0)
	for rows.Next() {
		var order entity.OrderOutputDTO
		if err = rows.Scan(&order.ID, &order.Item, &order.Amount); err != nil {
			return orders, err
		}
		orders = append(orders, &order)
	}

	return orders, nil
}

func (r *OrderPostgresRepository) CreateOrder(order *entity.Order) (*entity.OrderOutputDTO, error) {
	if order == nil {
		return nil, ErrInvalidEntity
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err := r.db.PrepareContext(ctx, `INSERT INTO orders(item, amount) VALUES($1, $2) RETURNING id`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(order.Item, order.Amount)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&order.ID); err != nil {
			return nil, err
		}
	}

	return &entity.OrderOutputDTO{
		ID:     order.ID,
		Item:   order.Item,
		Amount: order.Amount,
	}, nil
}

func (r *OrderPostgresRepository) GetOrderByID(id int) (*entity.OrderOutputDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := r.db.QueryRowContext(ctx, "SELECT * FROM orders WHERE id = $1", id)

	var data string
	if err := row.Scan(&data); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrOrderNotFound
		}
		return nil, err
	}

	order := &entity.OrderInputDTO{}
	if err := json.Unmarshal([]byte(data), order); err != nil {
		return nil, err
	}

	return &entity.OrderOutputDTO{
		ID:     order.ID,
		Item:   order.Item,
		Amount: order.Amount,
	}, nil
}

func (r *OrderPostgresRepository) UpdateOrder(id int, order *entity.OrderInputDTO) (*entity.OrderOutputDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err := r.db.PrepareContext(ctx, `UPDATE orders SET item = $1, amount = $2, WHERE id = $3`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(order.Item, order.Amount, id); err != nil {
		return nil, err
	}

	return &entity.OrderOutputDTO{
		ID:     order.ID,
		Item:   order.Item,
		Amount: order.Amount,
	}, nil
}

func (r *OrderPostgresRepository) DeleteOrder(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err := r.db.PrepareContext(ctx, `DELETE FROM orders WHERE id = $1 RETURNING id`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func NewOrderPostgresRepository() (entity.OrderRepository, error) {
	dataSourceName := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=%s",
		config.G.Db.User, config.G.Db.Dbname, config.G.Db.Password, config.G.Db.Host, config.G.Db.Port, config.G.Db.Sslmode)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(config.G.Db.MaxOpenConns)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance("file://internal/repository/migrations", "postgres", driver)
	if err != nil {
		return nil, err
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, err
	}

	return &OrderPostgresRepository{db: db}, nil
}
