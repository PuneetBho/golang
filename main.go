package main

import (
	"context"
	"flag"
	"go-mongo/repository"
	"go-mongo/services"
	"log"
	"net"

	sv "go-mongo/proto"

	"google.golang.org/grpc"
)

func main() {

	ctx := context.Background()
	// Connect to the database.
	var dbURL string
	var dbName string
	var serverPort string

	flag.StringVar(&dbURL, "url", "mongodb://localhost:27018", "")
	flag.StringVar(&dbName, "db name", "ownershubdb", "")
	flag.StringVar(&serverPort, "grpc server", "0.0.0.0:5001", "")
	flag.Parse()

	repo, err := repository.NewRepository(ctx, dbName, dbURL)
	if err != nil {
		log.Panic("database connect failed ", err)
	}

	log.Println("connected to DB")

	server := grpc.NewServer()

	sv.RegisterPlayerServiceServer(server, services.NewService(*repo))

	lis, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Panic("gRPC failed to start", err, "port", server)
		return // stop service.
	}
	log.Println("registered gRPC server")

	if err = server.Serve(lis); err != nil {
		log.Fatalf("Failed to server %v\n", err)
	}

	defer repo.Disconnect(ctx)

	defer server.GracefulStop()

}
