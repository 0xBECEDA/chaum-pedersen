package server

import (
	"context"
	"github.com/google/uuid"
	"log"
	v22 "nillion/api/v2"
	"nillion/internal/storage"
	"nillion/internal/utils"
	"nillion/internal/zkp"
)

type Server struct {
	v22.UnimplementedAuthServer
	verifier *zkp.Verifier
	db       storage.DB
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

func (s *Server) Register(_ context.Context, req *v22.RegisterRequest) (*v22.RegisterResponse, error) {
	y1, err := utils.ParseBigIntParam(req.GetY1(), "y1")
	if err != nil {
		return nil, err
	}

	y2, err := utils.ParseBigIntParam(req.GetY2(), "y2")
	if err != nil {
		return nil, err
	}

	s.db.RegisterUser(req.GetUser(), y1, y2)

	log.Printf("Register: registered user %v", req.GetUser())
	return &v22.RegisterResponse{}, nil
}

func (s *Server) CreateAuthenticationChallenge(_ context.Context, req *v22.AuthenticationChallengeRequest) (*v22.AuthenticationChallengeResponse, error) {
	// We use the google's widely used `uuid` pkg to generate the authID
	authID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	r1, err := utils.ParseBigIntParam(req.R1, "r1")
	if err != nil {
		return nil, err
	}

	r2, err := utils.ParseBigIntParam(req.R2, "r2")
	if err != nil {
		return nil, err
	}

	// Store the generated random value `c` and the `auth_id` in the authentication directory
	// for authentication verification process in the next step
	c := s.verifier.GenerateC()
	if err := s.db.AddAuthValues(authID.String(), req.GetUser(), r1, r2, c); err != nil {
		return nil, err
	}

	log.Printf("CreateAuthenticationChallenge: created challenge fot user %v", req.GetUser())
	log.Printf("CreateAuthenticationChallenge: user auth id %v", authID.String())

	return &v22.AuthenticationChallengeResponse{
		AuthId: authID.String(),
		C:      c.String(),
	}, nil
}

func (s *Server) VerifyAuthentication(_ context.Context, req *v22.AuthenticationAnswerRequest) (*v22.AuthenticationAnswerResponse, error) {
	authData, err := s.db.GetUserAuthData(req.GetAuthId())
	if err != nil {
		return nil, err
	}

	y1, y2, err := s.db.GetUserRegData(authData.GetUserID())
	if err != nil {
		return nil, err
	}

	sVal, err := utils.ParseBigIntParam(req.GetS(), "s")
	if err != nil {
		return nil, err
	}

	if err := s.verifier.Verify(sVal, authData.GetC(), authData.GetR1(), authData.GetR2(), y1, y2); err != nil {
		// return error if computed r1 and r2 are not same with expected r1 and r2
		return nil, err
	}

	log.Printf("VerifyAuthentication: user with auth id %v verified", req.GetAuthId())

	// If a valid proof is presented - then generate a sessionID and pass it as a response
	sessionID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &v22.AuthenticationAnswerResponse{SessionId: sessionID.String()}, nil
}
