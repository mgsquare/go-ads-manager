package ads

type Ad struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description,omitempty"`
	VideoURL     string `json:"video_url"`
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
	Duration     int    `json:"duration_seconds"`
	TargetURL    string `json:"target_url"`
}
