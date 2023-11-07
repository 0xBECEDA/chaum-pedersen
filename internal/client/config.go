package client

import (
	"errors"
	"math/big"
	"nillion/internal/utils"
	"os"
)

var (
	ErrEmptyPort        = errors.New("empty client port")
	ErrEmptyHost        = errors.New("empty client host")
	ErrEmptyUsername    = errors.New("empty username")
	ErrEmptySecretValue = errors.New("empty secret value")
)

type Config struct {
	Hostname    string
	Port        string
	Username    string
	SecretValue *big.Int
	G           *big.Int
	H           *big.Int
	P           *big.Int
}

func (c *Config) Load() error {
	host := os.Getenv("HOSTNAME")
	if host == "" {
		return ErrEmptyHost
	}
	c.Hostname = host

	port := os.Getenv("PORT")
	if port == "" {
		return ErrEmptyPort
	}
	c.Port = port

	username := os.Getenv("USERNAME")
	if username == "" {
		return ErrEmptyUsername
	}
	c.Username = username

	secret := os.Getenv("SECRET")
	if secret == "" {
		return ErrEmptySecretValue
	}

	secretNum, err := utils.ParseBigIntParam(secret, "secret")
	if err != nil {
		return err
	}
	c.SecretValue = secretNum

	g, h, p, err := utils.ReadProtocolValueFormEnv()
	if err != nil {
		return err
	}

	c.G = g
	c.H = h
	c.P = p
	return nil
}
