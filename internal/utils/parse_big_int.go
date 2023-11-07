package utils

import (
	"fmt"
	"math/big"
)

// ParseBigIntParam Parses a string and returns a pointer to the
// big.Int if successful
func ParseBigIntParam(str, param string) (*big.Int, error) {
	bigInt := new(big.Int)
	bigInt, valid := bigInt.SetString(str, 10)
	if !valid {
		return nil, fmt.Errorf("error parsing string %s to big.Int", param)
	}
	return bigInt, nil
}
