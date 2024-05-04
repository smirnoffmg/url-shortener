package repositories

import (
	entities "github.com/smirnoffmg/url-shortener/entities"
)

type Repository interface {
	Create(url, alias string) *entities.UrlRecord
	Get(alias string) *entities.UrlRecord
	List() []entities.UrlRecord
	IncrementVisits(alias string) error
}
