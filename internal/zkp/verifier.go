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

func (v *Verifier) GenerateC() *big.Int {
	return utils.GenerateRandomNonZeroValue()
}

// Verify verifies that expected and actual r1 and r2 values are equal
func (v *Verifier) Verify(s, c, r1, r2, y1, y2 *big.Int) error {
	// r1 = (g^s * y1 ^c) mod p
	// r2 = (h^s * y2 ^c) mod p

	gExp := big.NewInt(0).Exp(v.g, s, nil)
	y1Exp := big.NewInt(0).Exp(y1, c, nil)

	subR1Res := big.NewInt(0).Mul(gExp, y1Exp)
	actualR1 := big.NewInt(0).Mod(subR1Res, v.p)

	if r1.Cmp(actualR1) != 0 {
		return ErrR1IsNotSame
	}

	hExp := big.NewInt(0).Exp(v.h, s, nil)
	y2Exp := big.NewInt(0).Exp(y2, c, nil)

	subR2Res := big.NewInt(0).Mul(hExp, y2Exp)
	actualR2 := big.NewInt(0).Mod(subR2Res, v.p)

	if r2.Cmp(actualR2) != 0 {
		return ErrR2IsNotSame
	}
	return nil
}
