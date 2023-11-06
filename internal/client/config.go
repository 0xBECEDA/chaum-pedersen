package client

import (
	"errors"
	"math/big"
	"nillion/internal/utils"
	"os"
)

var (
	ErrEmptyPort = errors.New("empty client port")
	ErrEmptyHost = errors.New("empty client host")
)

type Config struct {
	Hostname string
	Port     string
	G        *big.Int
	H        *big.Int
	Q        *big.Int
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
	g, h, q, err := utils.ReadProtocolValueFormEnv()
	if err != nil {
		return err
	}

	c.G = g
	c.H = h
	c.Q = q
	return nil
}
