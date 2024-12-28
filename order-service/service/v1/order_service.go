package v1

import (
	"context"
	"errors"
	"log"
	"time"

	product "github.com/mayur-lomate-personal/grpc-order-product-app/product-service/grpc/api/v1"
	"google.golang.org/grpc" // For gRPC status codes
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	ErrProductNotFound   = errors.New("product not found")
	ErrInsufficientStock = errors.New("insufficient stock")
)

// Order represents the order entity
type Order struct {
	OrderID    int32
	Status     string
	TotalPrice float64
}

// OrderService handles business logic for orders
type OrderService struct {
	productClient product.ProductServiceClient // Injected ProductService gRPC client
}

// NewOrderService initializes the OrderService with a gRPC ProductService client
func NewOrderService(productConn *grpc.ClientConn) *OrderService {
	return &OrderService{
		productClient: product.NewProductServiceClient(productConn),
	}
}

// PlaceOrder processes an order
func (os *OrderService) PlaceOrder(orderID, productID, quantity int32, token string) (*Order, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("authorization", token))

	// Step 1: Get Product Details
	productResp, err := os.productClient.GetProductDetails(ctx, &product.GetProductRequest{
		ProductId: productID,
	})
	if err != nil {
		if statusErr, ok := status.FromError(err); ok {
			log.Printf("gRPC Error - Code: %v, Message: %v", statusErr.Code(), statusErr.Message())
		} else {
			log.Printf("Unknown Error: %v", err)
		}
		return nil, ErrProductNotFound
	}

	// Step 2: Check Stock Availability
	if productResp.Stock < quantity {
		return nil, ErrInsufficientStock
	}

	// Step 3: Update Stock
	_, err = os.productClient.UpdateStock(ctx, &product.UpdateStockRequest{
		ProductId: productID,
		Quantity:  -quantity, // Reduce stock
	})
	if err != nil {
		return nil, err
	}

	// Step 4: Calculate Total Price and Return Order Details
	totalPrice := productResp.Price * float64(quantity)
	order := &Order{
		OrderID:    orderID,
		Status:     "CONFIRMED",
		TotalPrice: totalPrice,
	}
	return order, nil
}
