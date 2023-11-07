package zkp

import (
	"math/big"
	"nillion/internal/utils"
)

type Verifier struct {
	g, h, p, q *big.Int
}

func NewVerifier(g, h, p, q *big.Int) *Verifier {
	return &Verifier{
		g: g,
		h: h,
		p: p,
		q: q,
	}
}

// GenerateC generates random non-zero value c
func (v *Verifier) GenerateC() *big.Int {
	return utils.GenerateRandomNonZeroValue()
}

// Verify verifies that expected and actual r1 and r2 values are equal
func (v *Verifier) Verify(s, c, r1, r2, y1, y2 *big.Int) error {

	// actual r1 = (g^s * y1 ^c) mod p
	gExp := new(big.Int).Exp(v.g, s, nil)
	y1Exp := new(big.Int).Exp(y1, c, nil)
	actualR1 := new(big.Int).Mul(gExp, y1Exp)
	actualR1.Mod(actualR1, v.p)

	if r1.Cmp(actualR1) != 0 {
		return ErrR1IsNotSame
	}

	// actual r2 = (h^s * y2 ^c) mod p
	hExp := new(big.Int).Exp(v.h, s, nil)
	y2Exp := new(big.Int).Exp(y2, c, nil)
	actualR2 := new(big.Int).Mul(hExp, y2Exp)
	actualR2.Mod(actualR2, v.p)

	if r2.Cmp(actualR2) != 0 {
		return ErrR2IsNotSame
	}
	return nil
}
