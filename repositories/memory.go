package repositories

import entities "github.com/smirnoffmg/url-shortener/entities"

type MemoryRepository struct {
	records map[string]entities.UrlRecord
}

func NewMemoryRepo() *MemoryRepository {
	return &MemoryRepository{
		records: make(map[string]entities.UrlRecord),
	}
}

func (m *MemoryRepository) Create(url, alias string) *entities.UrlRecord {
	newRecord := entities.UrlRecord{
		OriginalUrl: url,
		Alias:       alias,
	}
	m.records[newRecord.Alias] = newRecord

	return &newRecord
}

func (m *MemoryRepository) Get(alias string) *entities.UrlRecord {
	originalUrl, ok := m.records[alias]

	if !ok {
		return nil
	}

	return &originalUrl

}

func (m *MemoryRepository) List() []entities.UrlRecord {
	result := make([]entities.UrlRecord, len(m.records))

	return result
}

func (m *MemoryRepository) IncrementVisits(alias string) error {
	record := m.Get(alias)

	if record != nil {
		record.Visits++

		m.records[alias] = *record
	}

	return nil

}
