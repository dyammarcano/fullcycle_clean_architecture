package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/dyammarcano/fullcycle_clean_architecture/internal/domain"
	"github.com/dyammarcano/fullcycle_clean_architecture/pkg/parameters"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/inovacc/config"
	_ "github.com/lib/pq"
)

type OrderPostgresRepository struct {
	db *sql.DB
}

func (r *OrderPostgresRepository) ListOrders() ([]*domain.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, "SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			slog.Error(">>> Error closing rows: ", slog.String("error", err.Error()))
		}
	}(rows)

	orders := make([]*domain.Order, 0)
	for rows.Next() {
		var order domain.Order
		var amountTmp float64

		if err = rows.Scan(&order.ID, &order.Item, &amountTmp); err != nil {
			return orders, err
		}
		order.Amount = float32(amountTmp)

		orders = append(orders, &order)
	}

	return orders, nil
}

func (r *OrderPostgresRepository) CreateOrder(order *domain.Order) (*domain.Order, error) {
	if order == nil {
		return nil, fmt.Errorf("invalid entity")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err := r.db.PrepareContext(ctx, `INSERT INTO orders(item, amount) VALUES($1, $2) RETURNING id`)
	if err != nil {
		return nil, err
	}
	defer func(stmt *sql.Stmt) {
		if err := stmt.Close(); err != nil {
			slog.Error(">>> Error closing statement: ", slog.String("error", err.Error()))
		}
	}(stmt)

	rows, err := stmt.Query(order.Item, order.Amount)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			slog.Error(">>> Error closing rows: ", slog.String("error", err.Error()))
		}
	}(rows)

	for rows.Next() {
		if err = rows.Scan(&order.ID); err != nil {
			return nil, err
		}
	}

	return order, nil
}

func (r *OrderPostgresRepository) GetOrderByID(id int) (*domain.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := r.db.QueryRowContext(ctx, "SELECT id, item, amount FROM orders WHERE id = $1", id)

	order := &domain.Order{}
	var amountTmp float64
	if err := row.Scan(&order.ID, &order.Item, &amountTmp); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("order not found")
		}

		return nil, err
	}

	return order, nil
}

func (r *OrderPostgresRepository) UpdateOrder(id int, order *domain.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err := r.db.PrepareContext(ctx, `UPDATE orders SET item = $1, amount = $2 WHERE id = $3`)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		if err := stmt.Close(); err != nil {
			slog.Error(">>> Error closing statement: ", slog.String("error", err.Error()))
		}
	}(stmt)

	if _, err = stmt.Exec(order.Item, order.Amount, id); err != nil {
		return err
	}

	return nil
}

func (r *OrderPostgresRepository) DeleteOrder(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err := r.db.PrepareContext(ctx, `DELETE FROM orders WHERE id = $1`)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		if err := stmt.Close(); err != nil {
			slog.Error(">>> Error closing statement: ", slog.String("error", err.Error()))
		}
	}(stmt)

	if _, err = stmt.Exec(id); err != nil {
		return err
	}

	return nil
}

func NewOrderPostgresRepository() (domain.OrderRepository, error) {
	cfg, err := config.GetServiceConfig[*parameters.Service]()
	if err != nil {
		log.Fatalf("Failed to get service config: %v", err)
	}

	dataSourceName := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=%s",
		cfg.Database.User, cfg.Database.Name, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, "disable")

	// Open the database connection

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

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
