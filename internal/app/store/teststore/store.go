package teststore

import (
	"github.com/ArtemGontar/web-api/internal/app/model"
	"github.com/ArtemGontar/web-api/internal/app/store"
)

type Store struct {
	userRepository *UserRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	return &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}
}
