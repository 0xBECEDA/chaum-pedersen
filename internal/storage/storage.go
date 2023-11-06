package storage

import (
	"math/big"
	"sync"
)

type DB interface {
	RegisterUser(userId string, y1 *big.Int, y2 *big.Int)
	AddUserRValues(userId string, r1 *big.Int, r2 *big.Int) error
	AddUserSValue(userId string, s *big.Int) error
	AddGeneratedC(userId string, c *big.Int) error
	CleanUpUserSession(userId string) error
}

type userData struct {
	y1 *big.Int
	y2 *big.Int
	r1 *big.Int
	r2 *big.Int
	c  *big.Int
	s  *big.Int
}

type Storage struct {
	memoryDB map[string]*userData
	mut      *sync.Mutex
}

func NewStorage() *Storage {
	return &Storage{
		memoryDB: make(map[string]*userData),
		mut:      &sync.Mutex{},
	}
}

func (s *Storage) RegisterUser(userId string, y1 *big.Int, y2 *big.Int) {
	s.mut.Lock()
	defer s.mut.Unlock()

	s.memoryDB[userId] = &userData{
		y1: y1,
		y2: y2,
	}
}

func (s *Storage) AddUserRValues(userId string, r1 *big.Int, r2 *big.Int) error {
	s.mut.Lock()
	defer s.mut.Unlock()

	data, ok := s.memoryDB[userId]
	if !ok {
		return ErrUserDoNotExist
	}

	data.r1 = r1
	data.r2 = r2
	return nil
}

func (s *Storage) AddUserSValue(userId string, sValue *big.Int) error {
	s.mut.Lock()
	defer s.mut.Unlock()

	data, ok := s.memoryDB[userId]
	if !ok {
		return ErrUserDoNotExist
	}

	data.s = sValue
	return nil
}

func (s *Storage) AddGeneratedC(userId string, c *big.Int) error {
	s.mut.Lock()
	defer s.mut.Unlock()

	data, ok := s.memoryDB[userId]
	if !ok {
		return ErrUserDoNotExist
	}

	data.c = c
	return nil
}

func (s *Storage) CleanUpUserSession(userId string) error {
	s.mut.Lock()
	defer s.mut.Unlock()

	data, ok := s.memoryDB[userId]
	if !ok {
		return ErrUserDoNotExist
	}

	data.r1 = nil
	data.r2 = nil
	data.s = nil
	data.c = nil

	return nil
}
