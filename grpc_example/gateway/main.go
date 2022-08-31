package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	userpb "grpc_example/proto/gen/go"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions:   protojson.MarshalOptions{},
		UnmarshalOptions: protojson.UnmarshalOptions{},
	}))
	err := userpb.RegisterUserServiceHandlerFromEndpoint(
		ctx,
		mux,
		":8081",
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	)
	if err != nil {
		log.Fatalf("cannot register auth service: %v", err)
	}
	log.Fatal(http.ListenAndServe(":8082", mux))
}
