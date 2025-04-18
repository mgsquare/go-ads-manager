package clicks

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type Repository struct {
	DB          *sql.DB
	RedisClient redis.Client
}

func (repo *Repository) AddClicksData(click Click) error {

	if repo.isDBAvailable() {

		var exists bool
		err := repo.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM ads WHERE id = $1)", click.AdID).Scan(&exists)
		if err != nil {
			return fmt.Errorf("failed to check if ad exists: %w", err)
		}
		if !exists {
			return fmt.Errorf("ad with ID %d does not exist", click.AdID)

		}

		var count int
		err = repo.DB.QueryRow(`
			SELECT COUNT(*) 
			FROM clicks 
			WHERE user_id = $1 AND ad_id = $2 AND timestamp > NOW() - INTERVAL '30 seconds'
		`, click.UserID, click.AdID).Scan(&count)

		if err != nil {
			return fmt.Errorf("failed to check for duplicate click: %w", err)
		}

		if count > 0 {
			return fmt.Errorf("user %d has already clicked ad %d within the last 30 seconds", click.UserID, click.AdID)
		}

		query := `
            INSERT INTO clicks (ad_id, timestamp, ip, playback_position, user_id)
            VALUES ($1, $2, $3, $4, $5)
        `

		_, err = repo.DB.Exec(query, click.AdID, click.Timestamp.UTC(), click.IP, click.PlaybackPosition, click.UserID)
		if err != nil {
			return fmt.Errorf("failed to save click: %w", err)
		}

		return nil
	} else {
		return repo.cacheClickData(click)
	}
}

func (repo *Repository) isDBAvailable() bool {
	var result int
	err := repo.DB.QueryRow("SELECT 1").Scan(&result)
	return err == nil
}

func (repo *Repository) cacheClickData(click Click) error {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("click:%d", click.ID)

	clickData, err := json.Marshal(click)
	if err != nil {
		return fmt.Errorf("failed to marshal click data: %w", err)
	}

	err = repo.RedisClient.Set(ctx, cacheKey, clickData, 24*time.Hour).Err()
	if err != nil {
		return fmt.Errorf("failed to cache click data in Redis: %w", err)
	}

	return nil
}

func (repo *Repository) StartCachedClickProcessor() {
	ticker := time.NewTicker(60 * time.Second)
	go func() {
		for range ticker.C {
			if repo.isDBAvailable() {
				ctx := context.Background()
				pattern := "click:*"

				iter := repo.RedisClient.Scan(ctx, 0, pattern, 0).Iterator()
				for iter.Next(ctx) {
					key := iter.Val()
					val, err := repo.RedisClient.Get(ctx, key).Result()
					if err != nil {
						fmt.Printf("Failed to get data for key %s: %v\n", key, err)
						continue
					}

					var click Click
					err = json.Unmarshal([]byte(val), &click)
					if err != nil {
						fmt.Printf("Failed to unmarshal click data for key %s: %v\n", key, err)
						continue
					}

					err = repo.AddClicksData(click)
					if err != nil {
						fmt.Printf("Failed to insert click data for key %s: %v\n", key, err)
						continue
					}

					err = repo.RedisClient.Del(ctx, key).Err()
					if err != nil {
						fmt.Printf("Failed to delete key %s from Redis: %v\n", key, err)
					}
				}

				if err := iter.Err(); err != nil {
					fmt.Printf("Iterator error: %v\n", err)
				}
			} else {
				fmt.Println("DB is unavailable. Will retry after 60 seconds.")
			}
		}
	}()
}
