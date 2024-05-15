package repositories

import (
	entities "github.com/smirnoffmg/url-shortener/internal/entities"
	"gorm.io/gorm"
)

type IVisitsRepository interface {
	Create(visit *entities.Visit) error
	GetVisitsByAlias(alias string) ([]entities.Visit, error)
}

type VisitsRepository struct {
	db *gorm.DB
}

func NewVisitsRepository(db *gorm.DB) *VisitsRepository {
	db.AutoMigrate(&entities.Visit{})

	return &VisitsRepository{
		db: db,
	}
}

func (repo *VisitsRepository) Create(visit *entities.Visit) error {
	return repo.db.Create(visit).Error
}

func (repo *VisitsRepository) GetVisitsByAlias(alias string) ([]entities.Visit, error) {
	var visits []entities.Visit
	err := repo.db.Where("alias = ?", alias).Find(&visits).Error
	if err != nil {
		return nil, err
	}
	return visits, nil
}
