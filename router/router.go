package router

import (
	controller "SENAT_GROUPIE_TRACKER/controller"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(tmpl *template.Template) http.Handler {
	r := chi.NewRouter()

	// static files
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	h := controller.NewHandlers(tmpl)

	r.Get("/", h.IndexHandler)
	r.Get("/dashboard", h.DashboardHandler)
	r.Get("/api/clan", controller.ClanHandler) // si ces handlers existent, adaptez-les pour utiliser Handlers si besoin
	r.Get("/api/ip", controller.IpHandler)

	return r
}
