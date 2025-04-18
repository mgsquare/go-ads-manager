package analytics

import (
	"database/sql"
	"fmt"
	"time"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (repo *Repository) GetCTR(adID int, request AnalyticsRequest) (float64, error) {
	var clicks, impressions int
	var timeWindow time.Time

	switch request.Duration {
	case DurationMinute:
		timeWindow = time.Now().Add(-time.Minute)
	case DurationHour:
		timeWindow = time.Now().Add(-time.Hour)
	case DurationDay:
		timeWindow = time.Now().Add(-24 * time.Hour)
	case DurationWeek:
		timeWindow = time.Now().Add(-7 * 24 * time.Hour)
	case DurationMonth:
		timeWindow = time.Now().Add(-30 * 24 * time.Hour)
	default:
		if request.LastXMinutes > 0 {
			timeWindow = time.Now().Add(-time.Duration(request.LastXMinutes) * time.Minute)
		} else {
			return 0, fmt.Errorf("invalid custom range: %d", request.LastXMinutes)
		}
	}

	err := repo.DB.QueryRow(`
		SELECT COUNT(*) 
		FROM clicks 
		WHERE ad_id = $1 AND timestamp > $2
	`, adID, timeWindow).Scan(&clicks)
	if err != nil {
		return 0, fmt.Errorf("error fetching clicks: %w", err)
	}

	err = repo.DB.QueryRow(`
		SELECT COUNT(*) 
		FROM impressions 
		WHERE ad_id = $1 AND timestamp > $2
	`, adID, timeWindow).Scan(&impressions)
	if err != nil {
		return 0, fmt.Errorf("error fetching impressions: %w", err)
	}

	if impressions == 0 {
		return 0, nil
	}

	return (float64(clicks) / float64(impressions)) * 100, nil
}

func (repo *Repository) GetClicksByDuration(adID int, duration string) (int, error) {
	var query string

	switch duration {
	case "hour":
		query = `SELECT COUNT(*) FROM clicks WHERE ad_id = $1 AND timestamp >= NOW() - INTERVAL '1 hour'`
	case "day":
		query = `SELECT COUNT(*) FROM clicks WHERE ad_id = $1 AND timestamp >= NOW() - INTERVAL '1 day'`
	case "week":
		query = `SELECT COUNT(*) FROM clicks WHERE ad_id = $1 AND timestamp >= NOW() - INTERVAL '1 week'`
	case "month":
		query = `SELECT COUNT(*) FROM clicks WHERE ad_id = $1 AND timestamp >= NOW() - INTERVAL '1 month'`
	default:
		return 0, fmt.Errorf("unsupported duration: %s", duration)
	}

	var count int
	err := repo.DB.QueryRow(query, adID).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error fetching clicks for %s: %w", duration, err)
	}

	return count, nil
}

func (repo *Repository) GetClicksInLastXMinutes(adID int, minutes int) (int, error) {
	query := `SELECT COUNT(*) FROM clicks WHERE ad_id = $1 AND timestamp >= NOW() - ($2 || ' minutes')::interval`

	var count int
	err := repo.DB.QueryRow(query, adID, minutes).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error fetching clicks in last %d minutes: %w", minutes, err)
	}

	return count, nil
}
