package main

import (
	"context"
	"net/http"
	"wails-templ-hmtx-project/src/cmd/render"

	"github.com/go-chi/chi/v5"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Methods below are not bound to Wails lifecycle :

// getRoutes returns the application routes from the src directory
func (a *App) getRoutes() *chi.Mux {
	return render.GetRenderingMux()
}

// Middleware returns a NotFound middleware that enables the
// wails application to pass on HTTP calls to the backend handlers
// instead of serving static files
// See https://wails.io/docs/reference/options#assetserver
func (a *App) Middlewares() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		r := a.getRoutes()
		r.NotFound(next.ServeHTTP)
		return r
	}
}
