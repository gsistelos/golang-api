package main

import (
	"fmt"
	"log"
	"net"
	"os"

	v1 "github.com/gsistelos/golang-gRPC-API/gen/user/v1"
	"github.com/gsistelos/golang-gRPC-API/server"
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

	port := os.Getenv("8000")
	if port == "" {
		port = "8000"
	}

	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return fmt.Errorf("failed to listen to port %s", port)
	}
	defer l.Close()

	return grpcServer.Serve(l)
}
