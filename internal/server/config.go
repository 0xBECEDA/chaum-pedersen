package server

import (
	"errors"
	"math/big"
	"nillion/internal/utils"
	"os"
)

var (
	ErrEmptyPort = errors.New("empty server port")
)

type Config struct {
	Port string

	G *big.Int
	H *big.Int
	P *big.Int
}

func (c *Config) Load() error {
	port := os.Getenv("PORT")
	if port == "" {
		return ErrEmptyPort
	}
	c.Port = port

	g, h, p, err := utils.ReadProtocolValueFormEnv()
	if err != nil {
		return err
	}
	c.G = g
	c.H = h
	c.P = p
	return nil
}
