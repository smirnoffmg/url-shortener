package repositories

import (
	entities "github.com/smirnoffmg/url-shortener/internal/entities"
	"gorm.io/gorm"
)

type IUrlRecordsRepository interface {
	Create(urlRecord *entities.UrlRecord) error
	GetByAlias(alias string) (*entities.UrlRecord, error)
}

type UrlRecordsRepository struct {
	db *gorm.DB
}

func NewUrlRecordsRepository(db *gorm.DB) *UrlRecordsRepository {
	db.AutoMigrate(&entities.UrlRecord{})

	return &UrlRecordsRepository{
		db: db,
	}
}

func (repo *UrlRecordsRepository) Create(urlRecord *entities.UrlRecord) error {
	return repo.db.Create(urlRecord).Error
}

func (repo *UrlRecordsRepository) GetByAlias(alias string) (*entities.UrlRecord, error) {
	var urlRecord entities.UrlRecord
	err := repo.db.Where("alias = ?", alias).First(&urlRecord).Error
	if err != nil {
		return nil, err
	}
	return &urlRecord, nil
}
