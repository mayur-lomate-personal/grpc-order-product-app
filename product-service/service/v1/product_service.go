package v1

import (
	"errors"
	"fmt"
	"log"
	"os"
	model "product-service/model/v1"
	"sync"

	"github.com/knadh/koanf/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ProductService struct {
	db *gorm.DB
}

var (
	productServiceInstance *ProductService
	once                   sync.Once
)

// NewProductService initializes the ProductService
func NewProductService(k *koanf.Koanf) (*ProductService, error) {
	once.Do(func() {
		// Initialize productServiceInstance only once
		db, err := gorm.Open(postgres.Open(getDatabaseDSN(k)), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
			os.Exit(1) // Exit on failure to connect to the database
		}

		// Auto-migrate the Product model
		err = db.AutoMigrate(&model.Product{})
		if err != nil {
			log.Fatalf("Failed to migrate database: %v", err)
			os.Exit(1) // Exit on failure to migrate the database
		}

		// Seed initial data if the table is empty
		var count int64
		db.Model(&model.Product{}).Count(&count)
		if count == 0 {
			products := []model.Product{
				{ID: 1, Name: "Laptop", Description: "High-performance laptop", Price: 1200.50, Stock: 10},
				{ID: 2, Name: "Smartphone", Description: "Latest model smartphone", Price: 800.99, Stock: 25},
			}
			db.Create(&products)
		}

		productServiceInstance = &ProductService{db: db}
	})

	return productServiceInstance, nil
}

func getDatabaseDSN(k *koanf.Koanf) string {
	host := k.String("db_host")
	user := k.String("db_user")
	password := k.String("db_password")
	dbName := k.String("db_name")
	port := k.String("db_port")

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbName, port)
}

// GetProductByID retrieves product details
func (ps *ProductService) GetProductByID(productID int32) (*model.Product, error) {
	var product model.Product
	err := ps.db.First(&product, productID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("product not found")
	}
	return &product, err
}

// UpdateStock updates the stock for a product
func (ps *ProductService) UpdateStock(productID int32, quantity int32) (int32, error) {
	var product model.Product
	err := ps.db.First(&product, productID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, errors.New("product not found")
	}

	product.Stock += quantity
	err = ps.db.Save(&product).Error
	if err != nil {
		return 0, err
	}
	return product.Stock, nil
}
