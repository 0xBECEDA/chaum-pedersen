package utils

import (
	"math/big"
	"os"
)

const (
	defaultG = 4
	defaultH = 3
	defaultP = 23
)

func ReadProtocolValueFormEnv() (*big.Int, *big.Int, *big.Int, error) {
	var err error
	g := new(big.Int)
	gStr := os.Getenv("G_VALUE")
	if gStr == "" {
		g = big.NewInt(defaultG)
	} else {
		g, err = ParseBigIntParam(gStr, "g")
		if err != nil {
			return nil, nil, nil, err
		}
	}

	h := new(big.Int)
	hStr := os.Getenv("H_VALUE")
	if hStr == "" {
		h = big.NewInt(defaultH)
	} else {
		h, err = ParseBigIntParam(hStr, "h")
		if err != nil {
			return nil, nil, nil, err
		}
	}

	p := new(big.Int)
	pStr := os.Getenv("P_VALUE")
	if pStr == "" {
		p = big.NewInt(defaultP)
	} else {
		p, err = ParseBigIntParam(pStr, "p")
		if err != nil {
			return nil, nil, nil, err
		}
	}
	return g, h, p, nil
}
