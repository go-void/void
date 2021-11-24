// Package app combines multiple dependencies and starts them when the application is started
package app

import (
	"errors"

	"github.com/go-void/portal/pkg/server"

	"github.com/go-void/void/internal/config"
	"github.com/go-void/void/internal/router"
	"github.com/go-void/void/internal/router/handlers"
	"github.com/go-void/void/internal/store"
)

var (
	ErrAlreadyRunning = errors.New("app: already running")
)

type App struct {
	router *router.Router
	server *server.Server
	store  *store.Store

	isRunning bool
}

// New returns a new app
func New(path string) (*App, error) {
	cfg, err := config.Read(path)
	if err != nil {
		return nil, err
	}

	hdl := handlers.New()
	rtr := router.New(cfg.RouterOptions)
	rtr.AddRoutes(hdl)

	srv := server.New()
	srv.Configure(cfg.DNSOptions)

	str := store.New()

	return &App{
		router: rtr,
		server: srv,
		store:  str,
	}, nil
}

func (a *App) Run() error {
	if a.isRunning {
		return ErrAlreadyRunning
	}
	a.isRunning = true

	err := a.server.Run()
	if err != nil {
		return err
	}

	err = a.router.Run()
	if err != nil {
		return err
	}

	return nil
}
