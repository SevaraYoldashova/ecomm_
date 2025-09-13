package main

import (
	"ecommerce-backend/internal/auth"
	"ecommerce-backend/pkg/db"

	"ecommerce-backend/internal/handler"
	"ecommerce-backend/internal/repository"
	"ecommerce-backend/internal/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database, err := db.ConnectPostgres()
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	// repositories
	productRepo := repository.NewProductRepo(database)
	orderRepo := repository.NewOrderRepository(database)

	// services
	productService := service.NewProductService(productRepo)
	orderService := service.NewOrderService(orderRepo)

	// handlers
	productHandler := handler.NewProductHandler(productService)
	orderHandler := handler.NewOrderHandler(orderService)
	authHandler := handler.NewAuthHandler()

	// router
	r := mux.NewRouter()

	// auth
	r.HandleFunc("/login", authHandler.Login).Methods("POST")

	// products
	r.Handle("/products", auth.JWTMiddleware(http.HandlerFunc(productHandler.GetAll))).Methods("GET")
	r.Handle("/products/{id}", auth.JWTMiddleware(http.HandlerFunc(productHandler.GetByID))).Methods("GET")
	r.Handle("/products", auth.JWTMiddleware(http.HandlerFunc(productHandler.Create))).Methods("POST")
	r.Handle("/products/{id}", auth.JWTMiddleware(http.HandlerFunc(productHandler.Update))).Methods("PUT")
	r.Handle("/products/{id}", auth.JWTMiddleware(http.HandlerFunc(productHandler.Delete))).Methods("DELETE")

	// orders
	r.Handle("/orders", auth.JWTMiddleware(http.HandlerFunc(orderHandler.GetAll))).Methods("GET")
	r.Handle("/orders/{id}", auth.JWTMiddleware(http.HandlerFunc(orderHandler.GetByID))).Methods("GET")
	r.Handle("/orders", auth.JWTMiddleware(http.HandlerFunc(orderHandler.Create))).Methods("POST")

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", r)
}
