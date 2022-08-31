package main

import (
	"google.golang.org/grpc"
	userpb "grpc_example/proto/gen/go"
	user "grpc_example/userservice"
	"log"
	"net"
)

func main() {
	log.SetFlags(log.Lshortfile)

	lis, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &user.Service{})
	log.Fatal(s.Serve(lis))
}
