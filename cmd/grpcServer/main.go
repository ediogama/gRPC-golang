package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/ediogama/gRPC-golang/internal/database"
	"github.com/ediogama/gRPC-golang/internal/pb"
	"github.com/ediogama/gRPC-golang/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	categoryDB := database.NewCategory(db)
	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Create a new CategoryService
	categoryService := service.NewCategoryService(*categoryDB)

	// Register the CategoryService to the gRPC server
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	// Create a new gRPC listener
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("could not listen to :50051: %v", err)
	}

	// Start the gRPC server
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("could not start gRPC server: %v", err)
	}
}
