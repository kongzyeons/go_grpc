package main

import (
	"fmt"
	"log"
	"my-package/config"
	"my-package/repository"
	"my-package/services"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := config.LoadConfig()
	db := config.NewDataBaseSqlite(cfg.PathSqlite)

	userRepo := repository.NewUserRepo(db)
	userSrv := services.NewUserGrpcServer(userRepo)

	s := grpc.NewServer()
	lis, err := net.Listen("tcp", cfg.AppUrl)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	services.RegisterUserGrpcServer(s, userSrv)
	fmt.Println("grpc start server on port :", cfg.AppUrl)
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
