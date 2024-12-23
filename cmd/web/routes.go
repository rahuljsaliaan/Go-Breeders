package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(60 * time.Second))

	mux.Get("/", app.ShowHome)

	fileServer := http.FileServer(http.Dir("./static/"))

	// DESC: Strip the /static prefix and generate a handler function from the file server
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	// NOTE:  A very important reusability feature of the chi router is the ability to define routes with URL parameters. You can render any page by passing the page name as a URL parameter
	mux.Get("/{page}", app.ShowPage)

	return mux
}
