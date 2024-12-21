module grpc-order-product-app

go 1.23.4

require (
	github.com/golang-jwt/jwt/v4 v4.5.1 // JWT library for authentication
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.0 // gRPC Gateway for REST support
	google.golang.org/grpc v1.57.0 // gRPC library
	google.golang.org/protobuf v1.30.0 // Protocol Buffers
)

require golang.org/x/net v0.14.0 // indirect dependency for gRPC Gateway
