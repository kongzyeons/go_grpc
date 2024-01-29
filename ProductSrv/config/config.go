package config

import "os"

type Config struct {
	AppUrl     string
	PathSqlite string
}

func LoadConfig() Config {
	cfg := Config{
		AppUrl:     os.Getenv("APP_URL"),
		PathSqlite: os.Getenv("PATH_SQLITE"),
	}
	return cfg
}
