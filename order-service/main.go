package main

import (
	"context"
	"log"
	"net"
	"net/http"
	controller "order-service/controller/v1"
	JWTFilter "order-service/filter/v1"
	model "order-service/grpc/api/v1"
	service "order-service/service/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	// Connect to ProductService gRPC server
	productConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to ProductService: %v", err)
	}
	defer productConn.Close()

	// Initialize OrderService with ProductService client
	orderService := service.NewOrderService(productConn)

	// Initialize the OrderController and inject OrderService
	orderController := &controller.OrderController{
		Service: orderService,
	}

	// gRPC Server setup
	grpcLis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen on gRPC port 50053: %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(JWTFilter.UnaryInterceptor),
	)
	model.RegisterOrderServiceServer(grpcServer, orderController)

	// REST Gateway setup
	mux := runtime.NewServeMux()
	ctx := context.Background()
	err = model.RegisterOrderServiceHandlerServer(ctx, mux, orderController)
	if err != nil {
		log.Fatalf("Failed to register REST gateway: %v", err)
	}

	// Wrap REST Gateway with JWT Middleware
	httpServer := &http.Server{
		Addr:    ":8082", // HTTP port for REST API
		Handler: JWTFilter.HTTPMiddleware(mux),
	}

	// Start gRPC Server
	go func() {
		log.Println("Order Service running on gRPC port 50053")
		if err := grpcServer.Serve(grpcLis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// Start HTTP Server
	log.Println("Order Service running on REST port 8082")
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("Failed to serve REST: %v", err)
	}
}
