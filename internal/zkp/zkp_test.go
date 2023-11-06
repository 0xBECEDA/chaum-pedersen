package zkp

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"nillion/internal/utils"
	"testing"
)

func TestZKPChallenge(t *testing.T) {
	p := big.NewInt(23)

	g := big.NewInt(4)
	h := big.NewInt(3)

	q, err := utils.VerifyInitialValues(g, h, p)

	prover := NewProver(g, h, p, q)
	verifier := NewVerifier(g, h, p, q)

	y1, y2 := prover.CalculateYValues()
	k := prover.GenerateK()
	r1, r2 := prover.CalculateRValues(k)

	c := verifier.GenerateC()

	s := prover.ComputeS(c)

	err = verifier.Verify(s, c, r1, r2, y1, y2)
	assert.Empty(t, err)
}
