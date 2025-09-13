package service

import (
	"ecommerce-backend/internal/model"
	"ecommerce-backend/internal/repository"
)

type ProductService struct {
	Repo *repository.ProductRepo
}

func NewProductService(repo *repository.ProductRepo) *ProductService {
	return &ProductService{Repo: repo}
}

func (s *ProductService) GetAllProducts() ([]model.Product, error) {
	return s.Repo.GetAll()
}

func (s *ProductService) GetProduct(id int) (model.Product, error) {
	return s.Repo.GetByID(id)
}

func (s *ProductService) CreateProduct(p model.Product) error {
	return s.Repo.Create(p)
}

func (s *ProductService) UpdateProduct(p model.Product) error {
	return s.Repo.Update(p)
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.Repo.Delete(id)
}
