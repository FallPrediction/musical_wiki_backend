package service

import (
	"context"
	"encoding/json"
	"fmt"
	"musical_wiki/global"
	"musical_wiki/models"
	"musical_wiki/repository"
	"musical_wiki/request"
	"time"
)

type CreditService struct {
	repo repository.CreditRepository
}

func (service *CreditService) IndexByActorId(actorId string) ([]models.Credit, error) {
	key := fmt.Sprint("creditList:actorId=", actorId)
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	bytes, err := global.Redis.Get(ctx, key).Bytes()
	if err == nil {
		var cacheCredits []models.Credit
		err = json.Unmarshal(bytes, &cacheCredits)
		if err != nil {
			global.Logger.Warn("json unmarshal error", err)
		} else {
			return cacheCredits, nil
		}
	}
	cancel()

	credits, creditsErr := service.repo.IndexByActorId(actorId)
	if creditsErr == nil {
		bytes, err = json.Marshal(credits)
		if err != nil {
			global.Logger.Warn("json marshal error", err)
		} else {
			ctx, cancel = context.WithTimeout(context.Background(), 500*time.Millisecond)
			defer cancel()
			global.Redis.Set(ctx, key, bytes, 24*time.Hour)
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
	iter := global.Redis.Scan(ctx, 0, "creditList:actorId=*", 0).Iterator()
	for iter.Next(ctx) {
		global.Redis.Del(ctx, iter.Val())
	}
	if ctx.Err() == context.DeadlineExceeded {
		global.Logger.Warn("delCreditListCache timeout")
	}
}
