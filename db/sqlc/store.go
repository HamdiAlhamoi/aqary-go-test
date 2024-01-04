package db

import (
	"errors"

	"github.com/jackc/pgx/v4"
)

var (
	ErrNotFound  = errors.New("resource not found")
	ErrForbidden = errors.New("action forbidden")
)

type Store interface {
	Querier 
}

type ConduitStore struct {
	*Queries 
	db       *pgx.Conn
}

func NewConduitStore(db *pgx.Conn) Store {
	return &ConduitStore{
		db:      db,
		Queries: New(db),
	}
}

