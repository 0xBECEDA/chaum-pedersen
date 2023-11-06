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
	g := new(big.Int)
	gStr := os.Getenv("G_VALUE")
	if gStr == "" {
		g = big.NewInt(defaultG)
	} else {
		_, success := g.SetString(gStr, 10)
		if !success {
			return nil, nil, nil, ErrSetValue
		}
	}

	h := new(big.Int)
	hStr := os.Getenv("H_VALUE")
	if hStr == "" {
		h = big.NewInt(defaultH)
	} else {
		_, success := h.SetString(hStr, 10)
		if !success {
			return nil, nil, nil, ErrSetValue
		}
	}

	p := new(big.Int)
	pStr := os.Getenv("P_VALUE")
	if pStr == "" {
		p = big.NewInt(defaultP)
	} else {
		_, success := p.SetString(pStr, 10)
		if !success {
			return nil, nil, nil, ErrSetValue
		}
	}
	return g, h, p, nil
}
