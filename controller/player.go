package controller

import (
	"SENAT_GROUPIE_TRACKER/web"
	"encoding/json"
	"net/http"
)

var client = web.NewClashClient("TON_TOKEN_ICI")

func PlayerHandler(w http.ResponseWriter, r *http.Request) {
	tag := r.URL.Query().Get("tag")
	if tag == "" {
		http.Error(w, "missing tag", http.StatusBadRequest)
		return
	}

	data, err := client.GetPlayer(tag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
