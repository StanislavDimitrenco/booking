package main

import (
	"github.com/StanislavDimitrenko/booking/pkg/config"
	"github.com/StanislavDimitrenko/booking/pkg/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routs(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", http.HandlerFunc(handler.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handler.Repo.About))

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux

}
