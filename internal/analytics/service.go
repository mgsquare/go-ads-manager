package analytics

import (
	"fmt"
)

type Service struct {
	Repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) GetAnalytics(req AnalyticsRequest) (float64, error) {
	if !req.Type.IsValid() {
		return 0, fmt.Errorf("invalid analytics type: %s", req.Type)
	}

	switch req.Type {
	case TypeCTR:
		return s.Repo.GetCTR(req.AdID, req)

	case TypeClicks:
		if req.LastXMinutes > 0 {
			count, err := s.Repo.GetClicksInLastXMinutes(req.AdID, req.LastXMinutes)
			return float64(count), err
		} else if req.Duration.IsValid() {
			count, err := s.Repo.GetClicksByDuration(req.AdID, string(req.Duration))
			return float64(count), err
		} else {
			return 0, fmt.Errorf("no valid duration or last_x_minutes provided for clicks")
		}

	default:
		return 0, fmt.Errorf("unsupported analytics type: %s", req.Type)
	}
}
