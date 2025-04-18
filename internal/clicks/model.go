package clicks

import "time"

type Click struct {
	ID               int       `json:"id"`
	AdID             int       `json:"ad_id"`
	Timestamp        time.Time `json:"timestamp"`
	IP               string    `json:"ip"`
	PlaybackPosition float64   `json:"playback_position"`
	UserID           int       `json:"user_id"`
}
