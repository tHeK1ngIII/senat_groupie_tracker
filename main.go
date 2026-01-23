package main

import (
	"SENAT_GROUPIE_TRACKER/router"
	"log"
	"net/http"
	"os"
)

func main() {
	// initialiser les templates embarqués
	initTemplates()

	// passer tmpl au router
	r := router.SetupRoutes(tmpl)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Serveur lancé sur http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
