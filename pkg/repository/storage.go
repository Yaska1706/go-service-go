package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	)

type Storage interface {
	GetUser(userID int) (string, error)
	CreateUser(name string) (string, error)
}

func (s *storage) CreateUser(name string) (string, error) {
	panic("implement me")
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{db: db}
}

func (s *storage) GetUser(userID int) (string, error) {
	return "hello world", nil
}
