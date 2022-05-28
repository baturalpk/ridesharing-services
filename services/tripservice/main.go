package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/joho/godotenv"

	"github.com/baturalpk/ridesharing-services/services/tripservice/trip"
)

func main() {
	_ = godotenv.Load()
	port := mustGetenv("PORT")
	dburi := mustGetenv("MONGO_URI")
	dbname := mustGetenv("MONGO_DBNAME")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to create listener: %v", err)
	}

	if err := InitMongoConnection(dburi, dbname); err != nil {
		log.Fatalf("failed to init mongodb connection: %v", err)
	}

	repo := trip.NewRepository(
		GetCollection("trips"),
	)
	gsrv := grpc.NewServer()
	RegisterServer(gsrv, repo)
	reflection.Register(gsrv)

	log.Println("starting tripservice")
	if err := gsrv.Serve(lis); err != nil {
		CloseMongoConnection()
		log.Fatalf("failed to serve: %v", err)
	}
}

func mustGetenv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("cannot getenv: %s", key)
	}
	return val
}
