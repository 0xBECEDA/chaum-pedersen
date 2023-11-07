package utils

import (
	"math/big"
	"math/rand"
)

const maxNum = 100000

func GenerateRandomNonZeroValue() *big.Int {
	num := rand.Intn(maxNum)
	if num == 0 {
		return GenerateRandomNonZeroValue()
	}
	return big.NewInt(int64(num))
}
