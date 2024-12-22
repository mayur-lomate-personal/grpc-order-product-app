module order-service

go 1.23.4

require (
	github.com/golang-jwt/jwt/v4 v4.5.1 // JWT library for authentication
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.25.1 // gRPC Gateway for REST support
	google.golang.org/grpc v1.69.2 // gRPC library
	google.golang.org/protobuf v1.36.0 // Protocol Buffers
)

require golang.org/x/net v0.33.0 // indirect; indirect dependency for gRPC Gateway

require (
	github.com/mayur-lomate-personal/grpc-order-product-app/product-service v0.0.0
	google.golang.org/genproto/googleapis/api v0.0.0-20241219192143-6b3ec007d9bb
)

replace github.com/mayur-lomate-personal/grpc-order-product-app/product-service => ../product-service

require (
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241219192143-6b3ec007d9bb // indirect
)
