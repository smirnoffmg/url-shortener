package services

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	entities "github.com/smirnoffmg/url-shortener/internal/entities"
	repositories "github.com/smirnoffmg/url-shortener/internal/repositories"
)

const MinAliasLength = 8

var ctx = context.Background()

type UrlService struct {
	urlsRepo  repositories.IUrlRecordsRepository
	redis     *redis.Client
	urlLength int
}

func NewUrlService(repo repositories.IUrlRecordsRepository, urlLength int) *UrlService {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &UrlService{
		urlsRepo:  repo,
		redis:     rdb,
		urlLength: max(urlLength, MinAliasLength),
	}
}

func (svc *UrlService) Create(url string) (alias string, err error) {
	if err := UrlIsValid(url); err != nil {
		return "", err
	}

	alias = randomStr(svc.urlLength)

	if err := svc.urlsRepo.Create(&entities.UrlRecord{
		OriginalUrl: url,
		Alias:       alias,
	}); err != nil {
		log.Fatalf("Error creating url record: %v", err)
	}

	// store the alias in the cache
	svc.redis.Set(ctx, alias, url, 5*time.Minute)

	return alias, nil
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

	record, err := svc.urlsRepo.GetByAlias(alias)

	if err != nil {
		return nil
	}

	// store the alias in the cache
	svc.redis.Set(ctx, alias, record.OriginalUrl, 5*time.Minute)

	return record
}
