package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/go-sql-driver/mysql"
	postV1 "github.com/gsistelos/golang-api/gen/post/v1"
	userV1 "github.com/gsistelos/golang-api/gen/user/v1"
	"github.com/gsistelos/golang-api/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func getConnectionString() (string, error) {
	dbUser, ok := os.LookupEnv("MYSQL_USER")
	if !ok {
		dbUser = "root"
	}

	dbPassword, ok := os.LookupEnv("MYSQL_PASSWORD")
	if !ok {
		return "", fmt.Errorf("MYSQL_PASSWORD environment variable is required")
	}

	dbName, ok := os.LookupEnv("MYSQL_DATABASE")
	if !ok {
		dbName = "mysql"
	}

	return fmt.Sprintf("%s:%s@/%s?parseTime=true", dbUser, dbPassword, dbName), nil
}

func run() error {
	connStr, err := getConnectionString()
	if err != nil {
		return err
	}

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	srv := server.New(db)

	postV1.RegisterPostServiceServer(grpcServer, srv)
	userV1.RegisterUserServiceServer(grpcServer, srv)
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
