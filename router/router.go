package router

import (
	controller "SENAT_GROUPIE_TRACKER/controller"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes() http.Handler {
	r := chi.NewRouter()

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.Get("/", controller.IndexHandler)
	// r.Get("/api/player", api.PlayerHandler)
	r.Get("/dashboard", controller.DashboardHandler)
	r.Get("/api/clan", controller.ClanHandler)
	r.Get("/api/ip", controller.IpHandler)

	return r
}
