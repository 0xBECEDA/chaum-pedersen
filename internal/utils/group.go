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
		// 'a' is divisible by 'p' without reminder
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

func valueBelongsToGroup(v *big.Int, primeOrder *big.Int) (*big.Int, error) {
	modRes := new(big.Int)
	modRes.Mod(v, primeOrder)

	// The value should be in the residue field modulo p - it should have value from 0 to p-1
	if modRes.Cmp(primeOrder) > 0 {
		return nil, ErrValueNotInGroup
	}

	return smallFermatTheoremCheck(v, primeOrder)
}

// VerifyInitialValues checks if values 'g' and 'h' belong to group p.
// To pass the check successfully, you need to meet the conditions:
// 1. The 'p' value is prime number.
// 2. The values 'g' and 'h' are in the residue field modulo p.
// 3. We can find such 'q', when g^q mod p = 1 AND h^q mod p = 1
func VerifyInitialValues(g, h, p *big.Int) (*big.Int, error) {
	if !IsPrime(p) {
		return nil, ErrIsNotPrime
	}

	gOrder, err := valueBelongsToGroup(g, p)
	if err != nil {
		return nil, err
	}

	hOrder, err := valueBelongsToGroup(h, p)
	if err != nil {
		return nil, err
	}

	if gOrder.Cmp(hOrder) != 0 {
		return nil, ErrOrdersNotSame
	}
	return gOrder, nil
}
