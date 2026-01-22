package controller

import (
	"SENAT_GROUPIE_TRACKER/web"
	"encoding/json"
	"net/http"
)

var client = web.NewClashClient("eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6IjhjNzdjYjBmLTYwYjYtNGRkMi04YjJmLTcxMGQ0YmUwOTA5MSIsImlhdCI6MTc2OTA5MjIwNCwic3ViIjoiZGV2ZWxvcGVyL2IwOTFiMjliLTdjMDMtMWE2Mi01MzJmLTU1Nzc3YWY1OWEzMCIsInNjb3BlcyI6WyJyb3lhbGUiXSwibGltaXRzIjpbeyJ0aWVyIjoiZGV2ZWxvcGVyL3NpbHZlciIsInR5cGUiOiJ0aHJvdHRsaW5nIn0seyJjaWRycyI6WyI4NS4xNjkuODcuOTgiXSwidHlwZSI6ImNsaWVudCJ9XX0.UdQ50nZUkBcUAo_f_7fsTIwBIcDzT0EYs_uPpuJgK1ThIqob1P9R6MsDB0fEbuQF6ZfRbigI69fjHJAO-UfDUQ")

func ClanHandler(w http.ResponseWriter, r *http.Request) {
	tag := r.URL.Query().Get("tag")
	if tag == "" {
		http.Error(w, "missing clan tag", http.StatusBadRequest)
		return
	}

	data, err := client.GetClan(tag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
