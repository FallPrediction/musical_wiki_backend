package service

import (
	"fmt"
	"musical_wiki/models"
	"musical_wiki/repository"
	"musical_wiki/request"
	"musical_wiki/utils"

	"go.uber.org/zap"
)

type CreditService struct {
	repo   repository.CreditRepository
	logger *zap.SugaredLogger
	cache  utils.Cache
}

func (service *CreditService) IndexByActorId(actorId string) ([]models.Credit, error) {
	key := fmt.Sprint("creditList:actorId=", actorId)
	var cacheCredits []models.Credit
	cacheErr := service.cache.Get(key, &cacheCredits)
	if cacheErr == nil {
		return cacheCredits, nil
	}

	credits, creditsErr := service.repo.IndexByActorId(actorId)
	if creditsErr == nil {
		service.cache.Set(key, credits)
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
	service.cache.ScanAndDel("creditList:actorId=*")
}

func NewCreditService(repo repository.CreditRepository, logger *zap.SugaredLogger, cache utils.Cache) CreditService {
	return CreditService{repo: repo, logger: logger, cache: cache}
}
