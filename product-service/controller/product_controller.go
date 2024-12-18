package controller

import (
	"context"

	pb "github.com/mayur-lomate-personal/grpc-order-product-app/product-service/grpc/api/v1/product"
	"github.com/mayur-lomate-personal/grpc-order-product-app/product-service/service"
	"github.com/mayur-lomate-personal/grpc-order-product-app/product-service/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductController struct {
	pb.UnimplementedProductServiceServer
	Service *service.ProductService
}

// GetProductDetails retrieves product details by ID
func (pc *ProductController) GetProductDetails(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	if err := util.ValidateProductID(req.ProductId); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	product, err := pc.Service.GetProductByID(req.ProductId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "product not found: %v", err)
	}

	return &pb.GetProductResponse{
		ProductId:   product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}, nil
}

// UpdateStock updates the stock of a product
func (pc *ProductController) UpdateStock(ctx context.Context, req *pb.UpdateStockRequest) (*pb.UpdateStockResponse, error) {
	if err := util.ValidateProductID(req.ProductId); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	updatedStock, err := pc.Service.UpdateStock(req.ProductId, req.NewStock)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update stock: %v", err)
	}

	return &pb.UpdateStockResponse{
		ProductId:    req.ProductId,
		UpdatedStock: updatedStock,
		Message:      "Stock updated successfully",
	}, nil
}
