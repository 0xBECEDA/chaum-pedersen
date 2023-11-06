package zkp

import (
	"math/big"
	"nillion/internal/utils"
)

type Prover struct {
	x, g, h, p, q, k *big.Int
}

func NewProver(g, h, p, q *big.Int) *Prover {
	return &Prover{
		x: utils.GenerateRandomNonZeroValue(),
		g: g,
		h: h,
		p: p,
		q: q,
	}
}

func (p *Prover) GenerateK() *big.Int {
	p.k = utils.GenerateRandomNonZeroValue()
	return p.k
}

func (p *Prover) CalculateRValues(k *big.Int) (*big.Int, *big.Int) {
	r1, r2 := new(big.Int), new(big.Int)

	// r1 = g^k (mod p)
	// r2 = h^k (mod p)
	return r1.Exp(p.g, k, p.p), r2.Exp(p.h, k, p.p)
}

func (p *Prover) CalculateYValues() (*big.Int, *big.Int) {
	y1, y2 := new(big.Int), new(big.Int)

	// y1 = g^x (mod p)
	// y2 = h^x (mod p)
	return y1.Exp(p.g, p.x, p.p), y2.Exp(p.h, p.x, p.p)
}

func (p *Prover) ComputeS(c *big.Int) *big.Int {
	// s = (k - c * x) (mod q)

	res := big.NewInt(0).Mul(c, p.x)
	res = res.Sub(p.k, res)
	return big.NewInt(0).Mod(res, p.q)
}
