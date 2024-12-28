package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	controller "product-service/controller/v1"
	JWTFilter "product-service/filter/v1"
	productgrpc "product-service/grpc/api/v1"
	service "product-service/service/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"google.golang.org/grpc"
)

func main() {

	k := koanf.New(".")

	// Get the profile from the environment variable
	profile := os.Getenv("PROFILE")
	if profile == "" {
		profile = "development" // default to development if PROFILE is not set
	}

	// Construct the file path for the properties file
	filePath := fmt.Sprintf("resources/application-%s.yaml", profile)

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatalf("File not found: %s", filePath)
	}

	err := k.Load(file.Provider(filePath), yaml.Parser())
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
		os.Exit(1)
	}

	// Pass Koanf to the service constructor
	productService, err := service.NewProductService(k)
	if err != nil {
		log.Fatalf("Error initializing ProductService: %v", err)
		os.Exit(1)
	}

	// Initialize the ProductController and inject the ProductService
	productController := &controller.ProductController{
		Service: productService,
	}

	// gRPC Server
	grpcLis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		os.Exit(1)
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(JWTFilter.UnaryInterceptor),
	)
	productgrpc.RegisterProductServiceServer(grpcServer, productController)

	// REST Gateway
	mux := runtime.NewServeMux()
	ctx := context.Background()
	err = productgrpc.RegisterProductServiceHandlerServer(ctx, mux, productController)
	if err != nil {
		log.Fatalf("Failed to register REST gateway: %v", err)
		os.Exit(1)
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
			os.Exit(1)
		}
	}()
	log.Println("Product Service running on REST port 8081")
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("Failed to serve REST: %v", err)
	}
}
