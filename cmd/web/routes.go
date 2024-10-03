package main

import (
	"github.com/Max144/bookings/pkg/config"
	"github.com/Max144/bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/search-availability", handlers.Repo.SearchAvailability)
	mux.Get("/reservation", handlers.Repo.ReservationMake)
	mux.Get("/rooms/standard-room", handlers.Repo.StandardRoom)
	mux.Get("/rooms/deluxe-room", handlers.Repo.DeluxeRoom)
	mux.Get("/rooms/family-room", handlers.Repo.FamilyRoom)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
