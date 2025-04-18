package analytics

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) GetAnalytics(w http.ResponseWriter, r *http.Request) {
	var req AnalyticsRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.AdID == 0 || !req.Type.IsValid() {
		http.Error(w, "invalid ad_id or analytics type", http.StatusBadRequest)
		return
	}

	if req.Type == TypeClicks {
		if req.LastXMinutes > 0 && req.Duration.IsValid() {
			http.Error(w, "cannot provide both duration and custom_range for clicks", http.StatusBadRequest)
			return
		}

		if req.LastXMinutes == 0 && !req.Duration.IsValid() {
			http.Error(w, "must provide either duration or custome_range for clicks", http.StatusBadRequest)
			return
		}
	}

	result, err := h.Service.GetAnalytics(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"ad_id": req.AdID,
		"type":  req.Type,
		"value": result,
	})
}
