package utils

import (
	"math/big"
)

func IsPrime(n *big.Int) bool {
	return n.ProbablyPrime(10)
}

// Fermat's theorem: 'p' is a prime number and 'a' is an integer not divisible by 'p',
// then (a^p − 1) − 1 is divisible by 'p'.
// Therefore a^(p-1) mod p = 1
func smallFermatTheoremCheck(v *big.Int, primeOrder *big.Int) (*big.Int, error) {
	zero := big.NewInt(0)
	one := big.NewInt(1)

	rem := new(big.Int)
	if rem.Mod(v, primeOrder).Cmp(zero) == 0 {
		return nil, ErrValueNotInGroup
	}

	valuerOrder := new(big.Int)
	valuerOrder.Sub(primeOrder, one)

	expVal := new(big.Int)
	expVal.Exp(v, valuerOrder, nil)
	expVal.Sub(expVal, one)
	if expVal.Mod(expVal, primeOrder).Cmp(zero) == 0 {
		return valuerOrder, nil
	}

	return nil, ErrValueNotInGroup
}

// ValueBelongsToGroup checks if value 'v' belongs to prime group with order 'primeOrder'
func valueBelongsToGroup(v *big.Int, primeOrder *big.Int) (*big.Int, error) {
	modRes := new(big.Int)
	modRes.Mod(v, primeOrder)

	if modRes.Cmp(primeOrder) >= 0 {
		return nil, ErrValueNotInGroup
	}

	return smallFermatTheoremCheck(v, primeOrder)
}

func VerifyInitialValues(g, h, q *big.Int) (*big.Int, error) {
	if !IsPrime(q) {
		return nil, ErrIsNotPrime
	}

	gOrder, err := valueBelongsToGroup(g, q)
	if err != nil {
		return nil, err
	}

	hOrder, err := valueBelongsToGroup(h, q)
	if err != nil {
		return nil, err
	}

	if gOrder.Cmp(hOrder) != 0 {
		return nil, ErrOrdersNotSame
	}
	return gOrder, nil
}
