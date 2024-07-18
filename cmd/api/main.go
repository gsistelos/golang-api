package main

import (
	"fmt"
	"log"
	"net"
	"os"

	v1 "github.com/gsistelos/grpc-api/gen/user/v1"
	"github.com/gsistelos/grpc-api/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	grpcServer := grpc.NewServer()
	srv := server.New()

	v1.RegisterUserServiceServer(grpcServer, srv)
	reflection.Register(grpcServer)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return fmt.Errorf("failed to listen to port %s", port)
	}
	defer lis.Close()

	return grpcServer.Serve(lis)
}
