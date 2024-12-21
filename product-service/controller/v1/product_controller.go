package v1

import (
	"context"

	v1Product "product-service/grpc/api/v1"
	v1Service "product-service/service/v1"
	v1Util "product-service/util/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductController struct {
	v1Product.UnimplementedProductServiceServer
	Service *v1Service.ProductService
}

// GetProductDetails retrieves product details by ID
func (pc *ProductController) GetProductDetails(ctx context.Context, req *v1Product.GetProductRequest) (*v1Product.GetProductResponse, error) {
	if err := v1Util.ValidateProductID(req.ProductId); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	product, err := pc.Service.GetProductByID(req.ProductId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "product not found: %v", err)
	}

	return &v1Product.GetProductResponse{
		ProductId:   product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}, nil
}

// UpdateStock updates the stock of a product
func (pc *ProductController) UpdateStock(ctx context.Context, req *v1Product.UpdateStockRequest) (*v1Product.UpdateStockResponse, error) {
	if err := v1Util.ValidateProductID(req.ProductId); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	updatedStock, err := pc.Service.UpdateStock(req.ProductId, req.Quantity)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update stock: %v", err)
	}

	return &v1Product.UpdateStockResponse{
		ProductId: req.ProductId,
		NewStock:  updatedStock,
		Message:   "Stock updated successfully",
	}, nil
}
