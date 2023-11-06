package zkp

import "errors"

var (
	ErrR1IsNotSame = errors.New("expected r1 value doesn't match actual r1 value")
	ErrR2IsNotSame = errors.New("expected r2 value doesn't match actual r2 value")
)
