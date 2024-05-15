package services

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	entities "github.com/smirnoffmg/url-shortener/internal/entities"
	repositories "github.com/smirnoffmg/url-shortener/internal/repositories"
)

const MinAliasLength = 8

var ctx = context.Background()

type UrlService struct {
	repo      repositories.Repository
	redis     *redis.Client
	urlLength int
}

func NewUrlService(repo repositories.Repository, urlLength int) *UrlService {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &UrlService{
		repo:      repo,
		redis:     rdb,
		urlLength: max(urlLength, MinAliasLength),
	}
}

func (svc *UrlService) Create(url string) *entities.UrlRecord {
	if !UrlIsValid(url) {
		return nil
	}

	alias := randomStr(svc.urlLength)
	newRecord := svc.repo.Create(url, alias)
	return newRecord
}

func (svc *UrlService) Get(alias string) *entities.UrlRecord {
	// first check if the alias is in the cache

	val, err := svc.redis.Get(ctx, alias).Result()

	if err == nil {
		return &entities.UrlRecord{
			OriginalUrl: val,
			Alias:       alias,
		}
	}

	record := svc.repo.Get(alias)

	if record == nil {
		return nil
	}

	// store the alias in the cache
	svc.redis.Set(ctx, alias, record.OriginalUrl, 5*time.Minute)

	// increment the visits
	svc.repo.IncrementVisits(alias)

	return record

}
