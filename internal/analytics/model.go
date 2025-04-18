package analytics

type AnalyticsType string

const (
	TypeCTR    AnalyticsType = "ctr"
	TypeClicks AnalyticsType = "clicks"
)

func (t AnalyticsType) IsValid() bool {
	return t == TypeCTR || t == TypeClicks
}

type Duration string

const (
	DurationMinute Duration = "minute"
	DurationHour   Duration = "hour"
	DurationDay    Duration = "day"
	DurationWeek   Duration = "week"
	DurationMonth  Duration = "month"
)

func (d Duration) IsValid() bool {
	switch d {
	case DurationMinute, DurationHour, DurationDay, DurationWeek, DurationMonth:
		return true
	default:
		return false
	}
}

type AnalyticsRequest struct {
	AdID         int           `json:"ad_id"`
	Type         AnalyticsType `json:"type"`
	Duration     Duration      `json:"duration,omitempty"`
	LastXMinutes int           `json:"custom_range,omitempty"`
}
