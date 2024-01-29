package main

import (
	"fmt"
	"log"
	"my-package/config"
	"my-package/routes"

	_ "my-package/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title microservice-grpc
// @version 1.0
// @description Your API description
// @host localhost:8080
// @BasePath

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := config.LoadingConfig()
	// defer cfg.UserGrpc.Close()
	// defer cfg.ProductGrpc.Close()
	defer func() {
		if err := cfg.ProductGrpc.Close(); err != nil {
			fmt.Println("Error closing ProductGrpc:", err)
		}
	}()

	// Defer closing UserGrpc second (executed first due to reverse order)
	defer func() {
		if err := cfg.UserGrpc.Close(); err != nil {
			fmt.Println("Error closing UserGrpc:", err)
		}
	}()

	app := gin.New()
	routes.NewRouterApp(app, cfg)
	fmt.Println("Get start app port:", cfg.AppUrl)
	app.Run(cfg.AppUrl)

}
