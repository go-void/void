// Package store provides function to create, retrieve, update and delete data in a database
package store

import "github.com/jmoiron/sqlx"

type Store struct {
	Conn sqlx.DB
}

func New() *Store {
	return &Store{}
}
