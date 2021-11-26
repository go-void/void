// Package store provides function to create, retrieve, update and delete data in a database
package store

import (
	"errors"
	"fmt"

	"github.com/go-void/void/internal/config"
	"github.com/go-void/void/internal/constants"

	"github.com/jmoiron/sqlx"
)

var (
	ErrInvalidDatabaseType = errors.New("Invalid database type")
)

type Store struct {
	conn *sqlx.DB
}

func New(c config.StoreOptions) (*Store, error) {
	switch c.Backend {
	case "mysql":
		conn, err := sqlx.Open("mysql", DSN(c))
		if err != nil {
			return nil, err
		}

		return &Store{
			conn: conn,
		}, nil
	}

	return nil, ErrInvalidDatabaseType
}

func DSN(c config.StoreOptions) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
		constants.StoreParams,
	)
}
