package v1

import (
	"errors"
	v1model "product-service/model/v1"
	"sync"
)

type ProductService struct {
	productStore map[int32]*v1model.Product
	mu           sync.RWMutex
}

// NewProductService initializes the ProductService
func NewProductService() *ProductService {
	return &ProductService{
		productStore: map[int32]*v1model.Product{
			1: {ID: 1, Name: "Laptop", Description: "High-performance laptop", Price: 1200.50, Stock: 10},
			2: {ID: 2, Name: "Smartphone", Description: "Latest model smartphone", Price: 800.99, Stock: 25},
		},
	}
}

// GetProductByID retrieves product details
func (ps *ProductService) GetProductByID(productID int32) (*v1model.Product, error) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	product, exists := ps.productStore[productID]
	if !exists {
		return nil, errors.New("product not found")
	}

	return product, nil
}

// UpdateStock updates the stock for a product
func (ps *ProductService) UpdateStock(productID int32, newStock int32) (int32, error) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	product, exists := ps.productStore[productID]
	if !exists {
		return 0, errors.New("product not found")
	}

	product.Stock = newStock
	return product.Stock, nil
}
