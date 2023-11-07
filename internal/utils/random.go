package utils

import (
	"math/big"
	"math/rand"
)

const maxNum = 100000

// GenerateRandomNonZeroValue generates random non-zero value, which is not bigger than maxNum
// We limit the random number since it will be used for exponentiation with a big int.
// As we are dealing with numbers greater than int64, unrestricted exponentiation would require
// significant computational resources
func GenerateRandomNonZeroValue() *big.Int {
	num := rand.Intn(maxNum)
	if num == 0 {
		return GenerateRandomNonZeroValue()
	}
	return big.NewInt(int64(num))
}
