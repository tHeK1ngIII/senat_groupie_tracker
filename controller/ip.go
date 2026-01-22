package controller

import (
	"io"
	"net/http"
)

func IpHandler(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("https://api.ipify.org?format=text")
	if err != nil {
		http.Error(w, "Impossible de récupérer l'IP", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()
	ip, _ := io.ReadAll(res.Body)
	w.Write(ip)
}
