package services

import (
	"github.com/smirnoffmg/url-shortener/internal/entities"
	"github.com/smirnoffmg/url-shortener/internal/repositories"
)

type VisitsService struct {
	repo repositories.IVisitsRepository
}

func NewVisitsService(repo repositories.IVisitsRepository) *VisitsService {
	return &VisitsService{
		repo: repo,
	}
}

func (svc *VisitsService) SaveVisit(alias, ipAddr, userAgent string) error {
	visit := &entities.Visit{
		Alias:     alias,
		IpAddr:    ipAddr,
		UserAgent: userAgent,
	}

	return svc.repo.Create(visit)
}

func (svc *VisitsService) GetVisitsByAlias(alias string) ([]entities.Visit, error) {
	return svc.repo.GetVisitsByAlias(alias)
}
