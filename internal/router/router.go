// Package router provides a HTTP router / handler
package router

import (
	"fmt"

	"github.com/go-void/void/internal/config"
	"github.com/go-void/void/internal/router/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	echo    *echo.Echo
	handler *handlers.Handler
	config  config.RouterOptions
}

func New(c config.RouterOptions) *Router {
	e := echo.New()

	// Set debug mode off
	e.Debug = false

	// Hide startup message
	e.HideBanner = true

	// Add middlewares
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	return &Router{
		config: c,
		echo:   e,
	}
}

func (r *Router) AddRoutes(handle *handlers.Handler) {
	//// UNPROTECTED ROUTES ////
	// SPA
	r.echo.Use(middleware.Static(r.config.Path))

	// Prometheus metrics
	r.echo.GET("/metrics", handle.GetMetrics)

	//// PROTECTED ROUTES ////
	// API
	api := r.echo.Group("/api")

	// Domains
	api.DELETE("/domains/:id", handle.DeleteDomain)
	api.POST("/domains", handle.UpdateDomain)
	api.GET("/domains/:id", handle.GetDomain)
	api.GET("/domains", handle.GetDomains)
	api.PUT("/domains", handle.AddDomain)

	// STATS
	stats := r.echo.Group("/stats")
	stats.GET("/queries", handle.GetQueries)
	stats.GET("/queries/:id", handle.GetQuery)

	r.handler = handle
}

func (r *Router) Run() error {
	port := fmt.Sprintf(":%d", r.config.Port)
	return r.echo.Start(port)
}
