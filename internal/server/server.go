package server

import (
	"context"
	"nillion/internal/api"
	"nillion/internal/storage"
	"nillion/internal/zkp"
)

type Server struct {
	verifier *zkp.Verifier

	db storage.DB
}

func NewServer(
	verifier *zkp.Verifier,
	db storage.DB,
) *Server {
	return &Server{
		verifier: verifier,
		db:       db,
	}
}

func (s *Server) Register(context.Context, *api.RegisterRequest) (*api.RegisterResponse, error) {
	return nil, nil
}

func (s *Server) CreateAuthenticationChallenge(context.Context, *api.AuthenticationChallengeRequest) (*api.AuthenticationChallengeResponse, error) {
	return nil, nil
}

func (s *Server) VerifyAuthentication(context.Context, *api.AuthenticationAnswerRequest) (*api.AuthenticationAnswerResponse, error) {
	return nil, nil
}

func (s *Server) mustEmbedUnimplementedAuthServer() {}
