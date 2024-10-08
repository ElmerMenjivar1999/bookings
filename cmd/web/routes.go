package main

import (
	"github.com/ElmerMenjivar1999/bookings/pkg/config"
	"github.com/ElmerMenjivar1999/bookings/pkg/handlers"
	"net/http"

	//"github.com/bmizerany/pat"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()

	// mux.Get("/",http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about",http.HandlerFunc(handlers.Repo.About))
	
	mux := chi.NewRouter()

	//middlewares
	mux.Use(middleware.Recoverer)
	//mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	


	mux.Get("/",handlers.Repo.Home)
	mux.Get("/about",handlers.Repo.About)

	fileServer := http.FileServer(http.Dir("./static/"))

	mux.Handle("/static/*",http.StripPrefix("/static",fileServer))
	

	return mux
}