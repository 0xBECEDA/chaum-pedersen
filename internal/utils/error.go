package utils

import "errors"

var (
	ErrIsNotPrime      = errors.New("value is not prime")
	ErrValueNotInGroup = errors.New("value does not belong to group")

	ErrSetValue      = errors.New("error setting up value")
	ErrOrdersNotSame = errors.New("g and h orders are not same")
)
