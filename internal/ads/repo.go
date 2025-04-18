package ads

import (
	"database/sql"
	"fmt"
)

type Repository struct {
	DB *sql.DB
}

func (repo *Repository) GetAllAds() ([]Ad, error) {
	var ads []Ad

	rows, err := repo.DB.Query("SELECT id, title, description, video_url, target_url, duration FROM ads")

	if err != nil {
		return nil, fmt.Errorf("GetAllAds query failed: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var ad Ad
		err := rows.Scan(&ad.ID, &ad.Title, &ad.Description, &ad.VideoURL, &ad.TargetURL, &ad.Duration)
		if err != nil {
			return nil, fmt.Errorf("GetAllAds scan failed: %w", err)
		}
		ads = append(ads, ad)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetAllAds row iteration failed: %w", err)
	}

	return ads, nil

}
