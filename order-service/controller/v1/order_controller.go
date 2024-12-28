package v1

import (
	"context"
	"errors"
	model "order-service/grpc/api/v1"
	service "order-service/service/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type OrderController struct {
	model.UnimplementedOrderServiceServer
	Service *service.OrderService
}

// PlaceOrder handles the PlaceOrder RPC request
func (oc *OrderController) PlaceOrder(ctx context.Context, req *model.PlaceOrderRequest) (*model.PlaceOrderResponse, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	// Retrieve Authorization Header
	authHeader := md.Get("authorization")

	if len(authHeader) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is missing")
	}

	// Token should be "Bearer <token>"
	token := authHeader[0]

	// Validate input
	if req.OrderId <= 0 || req.ProductId <= 0 || req.Quantity <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid input: order_id, product_id, and quantity must be positive integers")
	}

	// Call service layer to process the order
	order, err := oc.Service.PlaceOrder(req.OrderId, req.ProductId, req.Quantity, token)
	if err != nil {
		if errors.Is(err, service.ErrProductNotFound) {
			return nil, status.Errorf(codes.NotFound, "product not found: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to place order: %v", err)
	}

	// Return successful response
	return &model.PlaceOrderResponse{
		OrderId:    order.OrderID,
		Status:     order.Status,
		TotalPrice: order.TotalPrice,
	}, nil
}
