package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/go-sql-driver/mysql"
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

func getDbCredentials() (string, string, string) {
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "root"
	}

	dbPassword := os.Getenv("DB_PASSWORD")

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "mysql"
	}

	return dbUser, dbPassword, dbName
}

func run() error {
	dbUser, dbPassword, dbName := getDbCredentials()
	if dbPassword == "" {
		return fmt.Errorf("DB_PASSWORD environment variable is required")
	}

	db, err := sql.Open("mysql", dbUser+":"+dbPassword+"@/"+dbName+"?parseTime=true")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	srv := server.New(db)

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
