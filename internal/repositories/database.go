package repositories

import (
	"log"
	"time"

	entities "github.com/smirnoffmg/url-shortener/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UrlDBRecord struct {
	gorm.Model
	OriginalUrl string
	Alias       string `gorm:"uniqueIndex"`
	Visits      int
}

type DBRepository struct {
	db *gorm.DB
}

func NewDBRepository(dbPath string) *DBRepository {
	log.Println("Connecting to database")
	dsn := "host=db user=myuser password=mypass dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf(err.Error())
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatalf(err.Error())
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	db.AutoMigrate(&UrlDBRecord{})

	return &DBRepository{
		db: db,
	}
}

func (repo *DBRepository) Create(url, alias string) *entities.UrlRecord {
	newRecord := &entities.UrlRecord{
		OriginalUrl: url,
		Alias:       alias,
	}

	repo.db.Create(&UrlDBRecord{
		OriginalUrl: url,
		Alias:       alias,
	})

	return newRecord
}

func (repo *DBRepository) Get(alias string) *entities.UrlRecord {
	var record UrlDBRecord

	if err := repo.db.Where("alias = ?", alias).First(&record).Error; err != nil {
		return nil
	}

	return &entities.UrlRecord{
		OriginalUrl: record.OriginalUrl,
		Alias:       record.Alias,
		Visits:      record.Visits,
	}

}

func (repo *DBRepository) List() (result []entities.UrlRecord) {
	rows, err := repo.db.Model(&UrlDBRecord{}).Rows()

	if err != nil {
		log.Fatalf(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var record UrlDBRecord
		repo.db.ScanRows(rows, &record)
		result = append(result, entities.UrlRecord{
			OriginalUrl: record.OriginalUrl,
			Alias:       record.Alias,
			Visits:      record.Visits,
		})
	}

	return
}

// IncrementVisits increments the number of visits for a given alias
func (repo *DBRepository) IncrementVisits(alias string) error {

	err := repo.db.Transaction(func(tx *gorm.DB) error {
		var record UrlDBRecord

		if err := tx.Where("alias = ?", alias).First(&record).Error; err != nil {
			return err
		}

		if err := tx.Model(&record).Update("visits", record.Visits+1).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil

}
