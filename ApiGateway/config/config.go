package config

import (
	"os"

	"google.golang.org/grpc"
)

type Config struct {
	AppUrl      string
	UserGrpc    *grpc.ClientConn
	ProductGrpc string
	OrderGrpc   string
}

func LoadingConfig() Config {
	cfg := Config{
		AppUrl:   os.Getenv("APP_URL"),
		UserGrpc: NewClientGrpc(os.Getenv("USER_GRPC")),
	}
	return cfg
}
