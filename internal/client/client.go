package client

import (
	"context"
	"log"
	v2 "nillion/api/v2"
	"nillion/internal/utils"
	"nillion/internal/zkp"
)

type Client struct {
	userName   string
	grpcClient v2.AuthClient
	prover     *zkp.Prover
}

func NewClient(
	userName string,
	grpcClient v2.AuthClient,
	prover *zkp.Prover) *Client {
	return &Client{
		userName:   userName,
		grpcClient: grpcClient,
		prover:     prover,
	}
}

type RegisterResp struct {
	Msg string `json:"msg"`
}

type LoginResp struct {
	SessionID string `json:"session_id"`
}

// Register registers the user with the given password and returns a message, if successful
func (c *Client) Register() (*RegisterResp, error) {
	y1, y2 := c.prover.CalculateYValues()

	log.Printf("y1 = %v, y2 = %v", y1.String(), y2.String())

	ctx := context.Background()
	_, err := c.grpcClient.Register(
		ctx,
		&v2.RegisterRequest{
			User: c.userName,
			Y1:   y1.String(),
			Y2:   y2.String(),
		},
	)

	if err != nil {
		return &RegisterResp{Msg: "registration is unsuccessful, try later"}, err
	}

	return &RegisterResp{Msg: "registration is successful"}, nil
}

// Login : Validates the login credentials using the Chaum-Pedersen Zero-Knowledge Proof
// protocol and returns a successful message for a valid login
func (c *Client) Login() (*LoginResp, error) {
	c.prover.GenerateK()
	r1, r2 := c.prover.CalculateRValues()

	// start auth challenge: send r1 and r2 values to server
	ctx := context.Background()
	authResp, err := c.grpcClient.CreateAuthenticationChallenge(
		ctx,
		&v2.AuthenticationChallengeRequest{
			User: c.userName,
			R1:   r1.String(),
			R2:   r2.String(),
		},
	)
	if err != nil {
		return nil, err
	}

	// get generated 'c' value
	authID := authResp.GetAuthId()
	cStr := authResp.GetC()
	cVal, err := utils.ParseBigIntParam(cStr, "c")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	s := c.prover.ComputeS(cVal)

	// verify step
	verifyRes, err := c.grpcClient.VerifyAuthentication(
		ctx,
		&v2.AuthenticationAnswerRequest{
			AuthId: authID,
			S:      s.String(),
		},
	)
	if err != nil {
		return nil, err
	}

	return &LoginResp{
		SessionID: verifyRes.SessionId,
	}, nil
}
