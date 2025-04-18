package clicks

type Service struct {
	Repo *Repository
}

func (s *Service) TrackClick(click Click) error {

	return s.Repo.AddClicksData(click)
}
