package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	v2 "nillion/api/v2"
	"nillion/internal/client"
	"nillion/internal/utils"
	"nillion/internal/zkp"
)

func main() {
	// load 'g', 'h', 'p' values from env or set up default values
	cfg := &client.Config{}
	if err := cfg.Load(); err != nil {
		log.Fatal(err)
	}

	// values should meet specific conditions and should be same with values on server
	q, err := utils.VerifyInitialValues(cfg.G, cfg.H, cfg.P)
	if err != nil {
		log.Fatal(err)
	}

	// create the gRPC client
	grpcClientOptions := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	addr := fmt.Sprintf("%v:%v", cfg.Hostname, cfg.Port)

	conn, err := grpc.Dial(addr, grpcClientOptions...)
	if err != nil {
		log.Fatalf("failed to dial server: %v", err)
	}
	grpcClient := v2.NewAuthClient(conn)

	prover := zkp.NewProver(cfg.SecretValue, cfg.G, cfg.H, cfg.P, q)
	cl := client.NewClient(cfg.Username, grpcClient, prover)

	// register client on server
	reg, err := cl.Register()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(reg.Msg)

	// login client on server
	logResp, err := cl.Login()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("successfull login, session id is %v", logResp.SessionID)
}
