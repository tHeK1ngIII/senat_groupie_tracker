package main

import (
	"SENAT_GROUPIE_TRACKER/router"
	"log"
	"net/http"
	"os"
)

func main() {
	r := router.SetupRoutes()
	port := os.Getenv("PORT") // Railway/Render te donnent ce port
	if port == "" {
		port = "8080" // fallback en local

	}
	log.Println("Serveur lancé sur http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
