package controller

import (
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Charge ton template principal
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, nil)
}
