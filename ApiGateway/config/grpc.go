package config

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClientGrpc(url string) *grpc.ClientConn {
	fmt.Println(url)
	creds := insecure.NewCredentials()
	cc, err := grpc.Dial(url, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	// defer cc.Close()
	return cc
}
