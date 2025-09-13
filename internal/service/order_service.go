package service

import (
	"ecommerce-backend/internal/model"
	"ecommerce-backend/internal/repository"
)

type OrderService struct {
	Repo *repository.OrderRepository
}

func NewOrderService(r *repository.OrderRepository) *OrderService {
	return &OrderService{Repo: r}
}

// Get all orders
func (s *OrderService) GetOrders() ([]model.Order, error) {
	return s.Repo.GetAll()
}

// Get one order by ID
func (s *OrderService) GetOrderByID(id int) (*model.Order, error) {
	return s.Repo.GetByID(id)
}

// Create new order
func (s *OrderService) CreateOrder(order *model.Order) error {
	return s.Repo.Create(order)
}
