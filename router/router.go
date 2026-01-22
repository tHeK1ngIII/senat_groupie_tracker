package router

import (
	"SENAT_GROUPIE_TRACKER/controller"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes() http.Handler {
	r := chi.NewRouter()

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.Get("/", controller.IndexHandler)
	r.Get("/api/player", controller.IndexHandler)
	r.Get("/dashboard", controller.DashboardHandler)

	return r
}
