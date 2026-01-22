package main

import (
	"SENAT_GROUPIE_TRACKER/router"
	"log"
	"net/http"
)

func main() {
	r := router.SetupRoutes()
	log.Println("Serveur lancé sur http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
