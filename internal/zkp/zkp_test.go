package zkp

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"nillion/internal/utils"
	"testing"
)

func TestZKPChallenge(t *testing.T) {
	t.Run("successful verification", func(t *testing.T) {
		const expectedQ = "22"

		p := big.NewInt(23)
		g := big.NewInt(4)
		h := big.NewInt(3)

		x := big.NewInt(24)

		// check that g and h belong to group p
		q, err := utils.VerifyInitialValues(g, h, p)
		assert.Empty(t, err)
		assert.Equal(t, expectedQ, q.String())

		prover := NewProver(x, g, h, p, q)
		verifier := NewVerifier(g, h, p, q)

		y1, y2 := prover.CalculateYValues()
		prover.GenerateK()

		// r1 = g^k mod p
		// r2 = h^k mod p
		r1, r2 := prover.CalculateRValues()

		c := verifier.GenerateC()

		// s = (k - c * x) mod q
		s := prover.ComputeS(c)

		// r1 = (g^s * y1 ^c) mod p
		// r2 = (h^s * y2 ^c) mod p
		err = verifier.Verify(s, c, r1, r2, y1, y2)
		assert.Empty(t, err)
	})

	t.Run("fail: g doesn't belong to group p", func(t *testing.T) {
		p := big.NewInt(23)
		g := big.NewInt(46)
		h := big.NewInt(3)

		// check that g and h belong to group p
		_, err := utils.VerifyInitialValues(g, h, p)
		assert.Equal(t, utils.ErrValueNotInGroup, err)
	})
}
