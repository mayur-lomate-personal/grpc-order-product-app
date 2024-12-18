module product-service

go 1.23.4

require (
	github.com/golang-jwt/jwt/v4 v4.5.1 // JWT library for authentication
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.24.0 // gRPC Gateway for REST support
	google.golang.org/grpc v1.69.0 // gRPC library
	google.golang.org/protobuf v1.36.0 // Protocol Buffers
)

require golang.org/x/net v0.30.0 // indirect dependency for gRPC Gateway

require (
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.20.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20241118233622-e639e219e697 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241118233622-e639e219e697 // indirect
)
