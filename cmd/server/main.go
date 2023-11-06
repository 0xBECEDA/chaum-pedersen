package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"nillion/internal/api"
	"nillion/internal/server"
	"nillion/internal/storage"
	"nillion/internal/utils"
	"nillion/internal/zkp"
)

func main() {
	cfg := &server.Config{}
	if err := cfg.Load(); err != nil {
		log.Fatal(err)
	}

	q, err := utils.VerifyInitialValues(cfg.G, cfg.H, cfg.P)
	if err != nil {
		log.Fatal(err)
	}

	verifier := zkp.NewVerifier(cfg.G, cfg.H, q, cfg.P)
	db := storage.NewStorage()
	grpcServ := grpc.NewServer()
	api.RegisterAuthServer(grpcServ, server.NewServer(verifier, db))

	listen, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := grpcServ.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
