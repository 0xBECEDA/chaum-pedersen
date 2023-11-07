package storage

import "errors"

var (
	ErrUserNotExists  = errors.New("user doesn't  exist")
	ErrAuthDoNotExist = errors.New("auth id does not exist")
)
