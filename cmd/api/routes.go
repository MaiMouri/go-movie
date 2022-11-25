package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// any time I have a variable type of application I have access to this function
func (app *application) routes() http.Handler {

	// create a router mux
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(app.enablesCORS)

	mux.Get("/", app.Home)

	mux.Get("/movies", app.AllMovies)
	mux.Get("/movies/{id}", app.GetMovie)
	return mux
}
