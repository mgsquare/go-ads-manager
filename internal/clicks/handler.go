package clicks

import (
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	Service *Service
}

func (h *Handler) TrackClickHandler(w http.ResponseWriter, r *http.Request) {

	var click Click
	if err := json.NewDecoder(r.Body).Decode(&click); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.Service.TrackClick(click); err != nil {
		http.Error(w, "Failed to record click", http.StatusInternalServerError)
		log.Print(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Click recorded successfully"}`))
}
