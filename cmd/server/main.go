package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"nillion/api/v2"
	"nillion/internal/server"
	"nillion/internal/storage"
	"nillion/internal/utils"
	"nillion/internal/zkp"
)

func main() {
	// load 'g', 'h', 'p' values from env or set up default values
	cfg := &server.Config{}
	if err := cfg.Load(); err != nil {
		log.Fatal(err)
	}

	// values should meet specific conditions and should be same with values on client
	q, err := utils.VerifyInitialValues(cfg.G, cfg.H, cfg.P)
	if err != nil {
		log.Fatal(err)
	}

	verifier := zkp.NewVerifier(cfg.G, cfg.H, cfg.P, q)
	db := storage.NewStorage() // in-memory storage

	// run grpc server
	grpcServ := grpc.NewServer()
	v2.RegisterAuthServer(grpcServ, server.NewServer(verifier, db))

	addr := ":" + cfg.Port
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server is running on %v", addr)
	if err := grpcServ.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
