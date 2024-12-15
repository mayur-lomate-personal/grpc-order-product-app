package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	pb "github.com/username/grpc-order-product-app/api/product"
	"github.com/username/grpc-order-product-app/internal/middleware"
	"github.com/username/grpc-order-product-app/internal/services/product"
)

func main() {
	// gRPC Server
	grpcLis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.UnaryInterceptor),
	)
	pb.RegisterProductServiceServer(grpcServer, &product.ProductServiceServer{})

	// REST Gateway
	mux := runtime.NewServeMux()
	ctx := context.Background()
	err = pb.RegisterProductServiceHandlerServer(ctx, mux, &product.ProductServiceServer{})
	if err != nil {
		log.Fatalf("Failed to register REST gateway: %v", err)
	}

	// Wrap REST Gateway with JWT Middleware
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: middleware.HTTPMiddleware(mux),
	}

	// Start gRPC and HTTP Servers
	go func() {
		log.Println("Product Service running on gRPC port 50052")
		if err := grpcServer.Serve(grpcLis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()
	log.Println("Product Service running on REST port 8080")
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("Failed to serve REST: %v", err)
	}
}
