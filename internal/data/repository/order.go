package repository

import (
	"context"
	"session-24/internal/data/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository struct {
	DB *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) Create(order *entity.Order) error {
	query := `
		INSERT INTO orders (email, total, created_at)
		VALUES ($1, $2, NOW())
	`
	_, err := r.DB.Exec(context.Background(), query, order.Email, order.Total)
	return err
}
