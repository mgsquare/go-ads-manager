package api

import (
	"net/http"

	"github.com/mgsquare/ads-tracking-backend/internal/ads"
	"github.com/mgsquare/ads-tracking-backend/internal/analytics"
	"github.com/mgsquare/ads-tracking-backend/internal/clicks"
)

func RegisterRoutes(adHandler *ads.Handler, clickHandler *clicks.Handler, analyticsHandler *analytics.Handler) {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok"}`))
	})

	http.HandleFunc("/ads", adHandler.GetAdsHandler)
	http.HandleFunc("/ads/click", clickHandler.TrackClickHandler)
	http.HandleFunc("/ads/analytics", analyticsHandler.GetAnalytics)
}
