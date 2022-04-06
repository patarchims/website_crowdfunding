package campaign

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignsDetailInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		campigns, err := s.repository.FindByUserID(userID)
		if err != nil {
			return campigns, err
		}
		return campigns, nil
	}

	campigns, err := s.repository.FindAll()
	if err != nil {
		return campigns, err
	}
	return campigns, nil
}

func (s *service) GetCampaignByID(input GetCampaignsDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(input.ID)

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
