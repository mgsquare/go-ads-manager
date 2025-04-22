package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/mgsquare/ads-tracking-backend/api"
	"github.com/mgsquare/ads-tracking-backend/internal/ads"
	"github.com/mgsquare/ads-tracking-backend/internal/analytics"
	"github.com/mgsquare/ads-tracking-backend/internal/clicks"
	database "github.com/mgsquare/ads-tracking-backend/internal/db"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}

	database.InitDB()
	//ads
	adRepo := &ads.Repository{DB: database.DB}
	adService := &ads.Service{Repo: adRepo}
	adHandler := &ads.Handler{Service: adService}
	//clicks
	clickRepo := &clicks.Repository{
		DB:          database.DB,
		RedisClient: *database.RedisClient,
	}
	clickService := &clicks.Service{Repo: clickRepo}
	clickHandler := &clicks.Handler{Service: clickService}
	//analytics
	analyticsRepo := &analytics.Repository{DB: database.DB}
	analyticsService := &analytics.Service{Repo: analyticsRepo}
	analyticsHandler := &analytics.Handler{Service: analyticsService}

	clickRepo.StartCachedClickProcessor()

	api.RegisterRoutes(adHandler, clickHandler, analyticsHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8008"
	}

	log.Printf("Server started on port %s", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
