package v1

import (
	"context"

	productgrpc "product-service/grpc/api/v1"
	service "product-service/service/v1"
	util "product-service/util/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductController struct {
	productgrpc.UnimplementedProductServiceServer
	Service *service.ProductService
}

// GetProductDetails retrieves product details by ID
func (pc *ProductController) GetProductDetails(ctx context.Context, req *productgrpc.GetProductRequest) (*productgrpc.GetProductResponse, error) {
	if err := util.ValidateProductID(req.ProductId); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	product, err := pc.Service.GetProductByID(req.ProductId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "product not found: %v", err)
	}

	return &productgrpc.GetProductResponse{
		ProductId:   product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}, nil
}

// UpdateStock updates the stock of a product
func (pc *ProductController) UpdateStock(ctx context.Context, req *productgrpc.UpdateStockRequest) (*productgrpc.UpdateStockResponse, error) {
	if err := util.ValidateProductID(req.ProductId); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	updatedStock, err := pc.Service.UpdateStock(req.ProductId, req.Quantity)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update stock: %v", err)
	}

	return &productgrpc.UpdateStockResponse{
		ProductId: req.ProductId,
		NewStock:  updatedStock,
		Message:   "Stock updated successfully",
	}, nil
}
