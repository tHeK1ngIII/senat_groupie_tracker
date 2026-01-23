package controller

import (
	"html/template"
	"log"
	"net/http"
)

type Handlers struct {
	Tmpl *template.Template
}

func NewHandlers(t *template.Template) *Handlers {
	return &Handlers{Tmpl: t}
}

func (h *Handlers) IndexHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{"Title": "Accueil"}
	if err := h.Tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
		log.Printf("template error: %v", err)
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
	}
}

func (h *Handlers) DashboardHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{"Title": "Dashboard"}
	if err := h.Tmpl.ExecuteTemplate(w, "dashboard.html", data); err != nil {
		log.Printf("template error: %v", err)
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
	}
}
