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
	api.POST("/domains/:id", handle.UpdateDomain)
	api.GET("/domains/:id", handle.GetDomain)
	api.GET("/domains", handle.GetDomains)
	api.PUT("/domains", handle.AddDomain)

	// Filters
	api.DELETE("/filters/:id", handle.DeleteFilter)
	api.POST("/filters/:id", handle.UpdateFilter)
	api.GET("/filters/:id", handle.GetFilter)
	api.GET("/filters", handle.GetFilters)
	api.PUT("/filters", handle.AddFilter)

	// Rules per filter
	filters := api.Group("/filters")
	filters.DELETE("/:fid/rules/:rid", handle.DeleteRule)
	filters.POST("/:fid/rules/:rid", handle.UpdateRule)
	filters.GET("/:fid/rules/:rid", handle.GetRule)
	filters.GET("/:fid/rules", handle.GetRules)
	filters.PUT("/:fid/rules", handle.AddRule)

	// STATS
	stats := r.echo.Group("/stats")
	stats.GET("/queries/:id", handle.GetQueryStat)
	stats.GET("/queries", handle.GetQueryStats)

	stats.GET("/filters/:id", handle.GetFilterStat)
	stats.GET("/filters", handle.GetFilterStats)

	r.handler = handle
}

func (r *Router) Run() error {
	port := fmt.Sprintf(":%d", r.config.Port)
	return r.echo.Start(port)
}
