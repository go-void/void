// Package app combines multiple dependencies and starts them when the application is started
package app

import (
	"github.com/go-void/portal/pkg/server"

	"github.com/go-void/void/internal/router"
	"github.com/go-void/void/internal/store"
)

type App struct {
	router *router.Router
	store  *store.Store
	server *server.Server
}

func New() (*App, error) {
	return nil, nil
}
