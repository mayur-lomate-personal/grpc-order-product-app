@echo off
REM Navigate to the root of the project

REM Define paths for Order Service
set PROTO_DIR=order-service\proto
set OUT_DIR=order-service\grpc

REM Generate gRPC and REST Gateway files for Order Service
protoc -I %PROTO_DIR% --go_out=paths=source_relative:%OUT_DIR% --go-grpc_out=paths=source_relative:%OUT_DIR% --grpc-gateway_out=paths=source_relative:%OUT_DIR% %PROTO_DIR%\api\v1\order.proto

REM Define paths for Product Service
set PROTO_DIR=product-service\proto
set OUT_DIR=product-service\grpc

REM Generate gRPC and REST Gateway files for Product Service
protoc -I %PROTO_DIR% --go_out=paths=source_relative:%OUT_DIR% --go-grpc_out=paths=source_relative:%OUT_DIR% --grpc-gateway_out=paths=source_relative:%OUT_DIR% %PROTO_DIR%\api\v1\product.proto

echo Proto files compiled successfully!
pause