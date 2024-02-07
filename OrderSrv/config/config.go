package config

import "os"

type Config struct {
	AppUrl   string
	MongoUrl string
}

func LoadConfig() Config {
	cfg := Config{
		AppUrl:   os.Getenv("APP_URL"),
		MongoUrl: os.Getenv("MONGO_URL"),
	}
	return cfg
}
