package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"musical_wiki/models"
	"musical_wiki/repository"
	"musical_wiki/request"
	"time"
)

type CreditService struct {
	repo   repository.CreditRepository
	logger *zap.SugaredLogger
	redis  *redis.Client
}

func (service *CreditService) IndexByActorId(actorId string) ([]models.Credit, error) {
	key := fmt.Sprint("creditList:actorId=", actorId)
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	bytes, err := service.redis.Get(ctx, key).Bytes()
	if err == nil {
		var cacheCredits []models.Credit
		err = json.Unmarshal(bytes, &cacheCredits)
		if err != nil {
			service.logger.Warn("json unmarshal error", err)
		} else {
			return cacheCredits, nil
		}
	}
	cancel()

	credits, creditsErr := service.repo.IndexByActorId(actorId)
	if creditsErr == nil {
		bytes, err = json.Marshal(credits)
		if err != nil {
			service.logger.Warn("json marshal error", err)
		} else {
			ctx, cancel = context.WithTimeout(context.Background(), 500*time.Millisecond)
			defer cancel()
			service.redis.Set(ctx, key, bytes, 24*time.Hour)
		}
	}
	return credits, creditsErr
}

func (service *CreditService) Store(request *request.Credit) (models.Credit, error) {
	credit := models.Credit{
		Time:      request.Time,
		Place:     request.Place,
		Character: request.Character,
		Musical:   request.Musical,
		ActorId:   request.ActorId,
	}
	service.delCreditListCache()
	return credit, service.repo.Store(&credit)
}

func (service *CreditService) Update(id string, request *request.Credit) error {
	credit := models.Credit{
		Time:      request.Time,
		Place:     request.Place,
		Character: request.Character,
		Musical:   request.Musical,
		ActorId:   request.ActorId,
	}
	service.delCreditListCache()
	return service.repo.Update(id, &credit)
}

func (service *CreditService) Destroy(id string) error {
	service.delCreditListCache()
	return service.repo.Destroy(id)
}

func (service *CreditService) delCreditListCache() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	iter := service.redis.Scan(ctx, 0, "creditList:actorId=*", 0).Iterator()
	for iter.Next(ctx) {
		service.redis.Del(ctx, iter.Val())
	}
	if ctx.Err() == context.DeadlineExceeded {
		service.logger.Warn("delCreditListCache timeout")
	}
}

func NewCreditService(repo repository.CreditRepository, logger *zap.SugaredLogger, redis *redis.Client) CreditService {
	return CreditService{repo: repo, logger: logger, redis: redis}
}
