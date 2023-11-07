package storage

import (
	"math/big"
	"sync"
)

type DB interface {
	RegisterUser(userId string, y1 *big.Int, y2 *big.Int)
	GetUserRegData(userId string) (*big.Int, *big.Int, error)
	GetUserAuthData(authID string) (*UserAuth, error)
	AddAuthValues(authID, userId string, r1 *big.Int, r2 *big.Int, c *big.Int) error
}

type Storage struct {
	regUsers  *usersDB
	authUsers *usersAuthDB
}

func NewStorage() *Storage {
	return &Storage{
		regUsers: &usersDB{
			memoryDB: map[string]*userData{},
			mut:      &sync.Mutex{},
		},
		authUsers: &usersAuthDB{
			memoryDB: map[string]*UserAuth{},
			mut:      &sync.Mutex{},
		},
	}
}

func (s *Storage) RegisterUser(userId string, y1 *big.Int, y2 *big.Int) {
	s.regUsers.mut.Lock()
	defer s.regUsers.mut.Unlock()

	s.regUsers.memoryDB[userId] = &userData{
		y1: y1,
		y2: y2,
	}
}

func (s *Storage) GetUserRegData(userId string) (*big.Int, *big.Int, error) {
	s.regUsers.mut.Lock()
	defer s.regUsers.mut.Unlock()

	data, ok := s.regUsers.memoryDB[userId]
	if !ok {
		return nil, nil, ErrUserNotExists
	}

	return data.y1, data.y2, nil
}

func (s *Storage) AddAuthValues(authID, userId string, r1, r2, c *big.Int) error {
	s.regUsers.mut.Lock()
	_, ok := s.regUsers.memoryDB[userId]
	if !ok {
		return ErrUserNotExists
	}
	s.regUsers.mut.Unlock()

	s.authUsers.mut.Lock()
	defer s.authUsers.mut.Unlock()

	s.authUsers.memoryDB[authID] = &UserAuth{
		userId: userId,
		r1:     r1,
		r2:     r2,
		c:      c,
	}
	return nil
}

func (s *Storage) GetUserAuthData(authID string) (*UserAuth, error) {
	s.authUsers.mut.Lock()
	defer s.authUsers.mut.Unlock()

	data, ok := s.authUsers.memoryDB[authID]
	if !ok {
		return nil, ErrAuthDoNotExist
	}

	return data, nil
}
