package repository

import (
	"database/sql"
	"ecommerce-backend/internal/model"
)

type ProductRepo struct {
	DB *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{DB: db}
}

func (r *ProductRepo) GetAll() ([]model.Product, error) {
	rows, err := r.DB.Query("SELECT id, name, description, price, quantity FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []model.Product{}
	for rows.Next() {
		var p model.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Quantity); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *ProductRepo) GetByID(id int) (model.Product, error) {
	var p model.Product
	err := r.DB.QueryRow("SELECT id, name, description, price, quantity FROM products WHERE id=$1", id).
		Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Quantity)
	return p, err
}

func (r *ProductRepo) Create(p model.Product) error {
	_, err := r.DB.Exec("INSERT INTO products(name, description, price, quantity) VALUES($1,$2,$3,$4)",
		p.Name, p.Description, p.Price, p.Quantity)
	return err
}

func (r *ProductRepo) Update(p model.Product) error {
	_, err := r.DB.Exec("UPDATE products SET name=$1, description=$2, price=$3, quantity=$4 WHERE id=$5",
		p.Name, p.Description, p.Price, p.Quantity, p.ID)
	return err
}

func (r *ProductRepo) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM products WHERE id=$1", id)
	return err
}
