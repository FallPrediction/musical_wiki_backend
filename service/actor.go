package service

import (
	"context"
	"encoding/json"
	"fmt"
	"musical_wiki/models"
	"musical_wiki/repository"
	"musical_wiki/request"
	"musical_wiki/utils"
	"strconv"
	"time"

	"go.uber.org/zap"
)

type ActorService struct {
	repo          repository.ActorRepository
	logger        *zap.SugaredLogger
	cache         utils.Cache
	creditService CreditService
	imageService  ImageService
}

func (service *ActorService) Index(currentPage int, perPage int) ([]models.Actor, int64, error) {
	key := fmt.Sprintf("actorsList:size=%v:currPage=%v", perPage, currentPage)
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	cache, err := service.cache.RedisClient.HGetAll(ctx, key).Result()
	if err == nil {
		var cacheActors []models.Actor
		err = json.Unmarshal([]byte(cache["actors"]), &cacheActors)
		if err == nil && len(cache) > 0 {
			service.logger.Warn("json unmarshal error", err)
		} else {
			count, _ := strconv.ParseInt(cache["count"], 10, 64)
			return cacheActors, count, nil
		}
	}

	actors, count, actorsErr := service.repo.Index(currentPage, perPage)
	if actorsErr == nil {
		bytes, err := json.Marshal(actors)
		if err != nil {
			service.logger.Warn("json marshal error", err)
		} else {
			ctx, cancel = context.WithTimeout(context.Background(), 500*time.Millisecond)
			defer cancel()
			service.cache.RedisClient.HSet(ctx, key, "actors", bytes, "count", count)
			service.cache.RedisClient.Expire(ctx, key, 24*time.Hour)
		}
	}
	return actors, count, actorsErr
}

func (service *ActorService) Show(id string) (models.Actor, error) {
	key := fmt.Sprintf("actor:%v", id)
	var cacheActor models.Actor
	cacheErr := service.cache.Get(key, &cacheActor)
	if cacheErr == nil {
		service.loadCredits(&cacheActor)
		service.loadGallery(&cacheActor)
		return cacheActor, nil
	}

	actor, actorErr := service.repo.Show(id)
	if actorErr == nil {
		service.cache.Set(key, actor)
	}
	service.loadCredits(&actor)
	service.loadGallery(&actor)
	return actor, actorErr
}

func (service *ActorService) Store(request *request.Actor) (models.Actor, error) {
	actor := models.Actor{
		Name:           request.Name,
		TranslatedName: request.TranslatedName,
		NickName:       request.NickName,
		Nationality:    request.Nationality,
		Born:           request.Born,
		Content:        request.Content,
		Socials:        request.Socials,
	}
	return actor, service.repo.Store(&actor)
}

func (service *ActorService) Update(id string, request *request.Actor) error {
	actor := models.Actor{
		Name:           request.Name,
		TranslatedName: request.TranslatedName,
		NickName:       request.NickName,
		Nationality:    request.Nationality,
		Born:           request.Born,
		Content:        request.Content,
		Socials:        request.Socials,
	}

	service.delActorCache(id)
	service.delActorsListCache()
	return service.repo.Update(id, &actor)
}

func (service *ActorService) Destroy(id string) error {
	service.delActorCache(id)
	service.delActorsListCache()
	return service.repo.Destroy(id)
}

func (service *ActorService) delActorCache(id string) {
	service.cache.Del(fmt.Sprintf("actor:%v", id))
}

func (service *ActorService) delActorsListCache() {
	service.cache.ScanAndDel("actorsList:size=*:currPage=*")
}

func (service *ActorService) loadCredits(actor *models.Actor) {
	credits, creditsErr := service.creditService.IndexByActorId(strconv.Itoa(int(actor.Id)))
	if creditsErr == nil {
		actor.Credits = credits
	}
}

func (service *ActorService) loadGallery(actor *models.Actor) {
	gallery, galleryErr := service.imageService.repo.IndexGallery(fmt.Sprint(actor.Id))
	if galleryErr == nil {
		actor.Gallery = gallery
	}
}

func NewActorService(repo repository.ActorRepository, logger *zap.SugaredLogger, cache utils.Cache, creditService CreditService, imageService ImageService) ActorService {
	return ActorService{repo: repo, logger: logger, cache: cache, creditService: creditService, imageService: imageService}
}
