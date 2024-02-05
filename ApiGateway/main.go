package main

import (
	"fmt"
	"log"
	"my-package/config"
	_ "my-package/docs"
	"my-package/routes"

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
	defer cfg.UserGrpc.Close()
	defer cfg.ProductGrpc.Close()

	app := gin.New()
	routes.NewRouterApp(app, cfg)
	fmt.Println("Get start app port:", cfg.AppUrl)
	app.Run(cfg.AppUrl)

}
