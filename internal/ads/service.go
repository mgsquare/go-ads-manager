package ads

type Service struct {
	Repo *Repository
}

func (s *Service) GetAllAds() ([]Ad, error) {
	return s.Repo.GetAllAds()
}
