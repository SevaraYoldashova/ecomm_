package repository

import (
	"database/sql"
	"ecommerce-backend/internal/model"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

// Fetch all orders
func (r *OrderRepository) GetAll() ([]model.Order, error) {
	rows, err := r.DB.Query("SELECT id, product_id, quantity, total FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []model.Order
	for rows.Next() {
		var o model.Order
		if err := rows.Scan(&o.ID, &o.ProductID, &o.Quantity, &o.Total); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}

	return orders, nil
}

// Fetch one order by ID
func (r *OrderRepository) GetByID(id int) (*model.Order, error) {
	var o model.Order
	err := r.DB.QueryRow("SELECT id, product_id, quantity, total FROM orders WHERE id = $1", id).
		Scan(&o.ID, &o.ProductID, &o.Quantity, &o.Total)
	if err != nil {
		return nil, err
	}
	return &o, nil
}

// Insert a new order
func (r *OrderRepository) Create(order *model.Order) error {
	err := r.DB.QueryRow(
		"INSERT INTO orders (product_id, quantity, total) VALUES ($1, $2, $3) RETURNING id",
		order.ProductID, order.Quantity, order.Total,
	).Scan(&order.ID)
	return err
}
