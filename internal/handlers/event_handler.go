package handlers

import (
	"encoding/json"
	"net/http"
	"study-event-api/internal/db"
)

func EventHandler(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		events, err := queries.ListEvents(r.Context())
		if err != nil {
			http.Error(w, "Unable to fetch events", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(events)
	}
}
