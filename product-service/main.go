package main

import (
	"context"
	"log"
	"net"
	"net/http"
	v1Controller "product-service/controller/v1"
	JWTFilter "product-service/filter/v1"
	v1Product "product-service/grpc/api/v1"
	v1Service "product-service/service/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {

	productService := v1Service.NewProductService()

	// Initialize the ProductController and inject the ProductService
	productController := &v1Controller.ProductController{
		Service: productService,
	}

	// gRPC Server
	grpcLis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(JWTFilter.UnaryInterceptor),
	)
	v1Product.RegisterProductServiceServer(grpcServer, productController)

	// REST Gateway
	mux := runtime.NewServeMux()
	ctx := context.Background()
	err = v1Product.RegisterProductServiceHandlerServer(ctx, mux, productController)
	if err != nil {
		log.Fatalf("Failed to register REST gateway: %v", err)
	}

	// Wrap REST Gateway with JWT Middleware
	httpServer := &http.Server{
		Addr:    ":8081",
		Handler: JWTFilter.HTTPMiddleware(mux),
	}

	// Start gRPC and HTTP Servers
	go func() {
		log.Println("Product Service running on gRPC port 50052")
		if err := grpcServer.Serve(grpcLis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()
	log.Println("Product Service running on REST port 8081")
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("Failed to serve REST: %v", err)
	}
}
