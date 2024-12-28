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
	gorm.io/driver/postgres v1.5.11
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.5.5 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/knadh/koanf/providers/file v1.1.2
	github.com/knadh/koanf/parsers/file v1.1.2
	github.com/knadh/koanf/v2 v2.1.2
	golang.org/x/crypto v0.28.0 // indirect
	golang.org/x/sync v0.9.0 // indirect
	gorm.io/gorm v1.25.10 // indirect
)

require (
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require (
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/knadh/koanf/parsers/yaml v0.1.0
)
