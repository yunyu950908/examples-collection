package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	userpb "grpc_example/proto/gen/go"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cannot connect server: %v", err)
	}
	usClient := userpb.NewUserServiceClient(conn)
	resp, err := usClient.GetUser(context.Background(), &userpb.GetUserRequest{Id: "user9527"})
	if err != nil {
		log.Fatalf("cannot call GetUser: %v", err)
	}
	fmt.Println(resp)
}
