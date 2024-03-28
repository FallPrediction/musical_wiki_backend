package service

import (
	"context"
	"encoding/json"
	"fmt"
	"musical_wiki/models"
	"musical_wiki/repository"
	"musical_wiki/request"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type ActorService struct {
	repo          repository.ActorRepository
	logger        *zap.SugaredLogger
	redis         *redis.Client
	creditService CreditService
	imageService  ImageService
}

func (service *ActorService) Index(currentPage int, perPage int) ([]models.Actor, int64, error) {
	key := fmt.Sprintf("actorsList:size=%v:currPage=%v", perPage, currentPage)
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	cache, err := service.redis.HGetAll(ctx, key).Result()
	if err == nil {
		var cacheActors []models.Actor
		err = json.Unmarshal([]byte(cache["actors"]), &cacheActors)
		if err != nil {
			service.logger.Warn("json unmarshal error", err)
		} else {
			count, _ := strconv.ParseInt(cache["count"], 10, 64)
			return cacheActors, count, nil
		}
	}
	cancel()

	actors, count, actorsErr := service.repo.Index(currentPage, perPage)
	if actorsErr == nil {
		bytes, err := json.Marshal(actors)
		if err != nil {
			service.logger.Warn("json marshal error", err)
		} else {
			ctx, cancel = context.WithTimeout(context.Background(), 500*time.Millisecond)
			defer cancel()
			service.redis.HSet(ctx, key, "actors", bytes, "count", count)
			service.redis.Expire(ctx, key, 24*time.Hour)
		}
	}
	return actors, count, actorsErr
}

func (service *ActorService) Show(id string) (models.Actor, error) {
	key := fmt.Sprintf("actor:%v", id)
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	bytes, err := service.redis.Get(ctx, key).Bytes()
	if err == nil {
		var cacheActor models.Actor
		err = json.Unmarshal(bytes, &cacheActor)
		if err != nil {
			service.logger.Warn("json unmarshal error", err)
		} else {
			service.loadCredits(&cacheActor)
			service.loadGallery(&cacheActor)
			return cacheActor, nil
		}
	}
	cancel()

	actor, actorErr := service.repo.Show(id)
	if actorErr == nil {
		bytes, err = json.Marshal(actor)
		if err != nil {
			service.logger.Warn("json marshal error", err)
		} else {
			ctx, cancel = context.WithTimeout(context.Background(), 500*time.Millisecond)
			defer cancel()
			service.redis.Set(ctx, key, bytes, 24*time.Hour)
		}
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	key := fmt.Sprintf("actor:%v", id)
	service.redis.Del(ctx, key)
	if ctx.Err() == context.DeadlineExceeded {
		service.logger.Warn("delActorCache timeout", key)
	}
}

func (service *ActorService) delActorsListCache() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	iter := service.redis.Scan(ctx, 0, "actorsList:size=*:currPage=*", 0).Iterator()
	for iter.Next(ctx) {
		service.redis.Del(ctx, iter.Val())
	}
	if ctx.Err() == context.DeadlineExceeded {
		service.logger.Warn("delActorsListCache timeout")
	}
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

func NewActorService(repo repository.ActorRepository, logger *zap.SugaredLogger, redis *redis.Client, creditService CreditService, imageService ImageService) ActorService {
	return ActorService{repo: repo, logger: logger, redis: redis, creditService: creditService, imageService: imageService}
}
