package storage

import (
	"math/big"
	"sync"
)

type userData struct {
	y1 *big.Int
	y2 *big.Int
}

type UserAuth struct {
	userId string
	r1     *big.Int
	r2     *big.Int
	c      *big.Int
}

func (ua *UserAuth) GetUserID() string {
	return ua.userId
}

func (ua *UserAuth) GetR1() *big.Int {
	return ua.r1
}

func (ua *UserAuth) GetR2() *big.Int {
	return ua.r2
}

func (ua *UserAuth) GetC() *big.Int {
	return ua.c
}

type usersDB struct {
	memoryDB map[string]*userData
	mut      *sync.Mutex
}

type usersAuthDB struct {
	memoryDB map[string]*UserAuth
	mut      *sync.Mutex
}
