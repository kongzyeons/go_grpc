package config

import (
	"os"

	"google.golang.org/grpc"
)

type Config struct {
	AppUrl      string
	UserGrpc    *grpc.ClientConn
	ProductGrpc *grpc.ClientConn
	OrderGrpc   *grpc.ClientConn
}

func LoadingConfig() Config {
	cfg := Config{
		AppUrl:      os.Getenv("APP_URL"),
		UserGrpc:    NewClientGrpc(os.Getenv("USER_GRPC")),
		ProductGrpc: NewClientGrpc(os.Getenv("PRODUCT_GRPC")),
		OrderGrpc:   NewClientGrpc(os.Getenv("ORDER_GRPC")),
	}
	return cfg
}
