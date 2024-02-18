package service

import (
	"context"
	"encoding/json"
	"fmt"
	"musical_wiki/global"
	"musical_wiki/models"
	"musical_wiki/repository"
	"musical_wiki/request"
	"strconv"
	"time"
)

type ActorService struct {
	repo repository.ActorRepository
}

func (service *ActorService) Index(currentPage int, perPage int) ([]models.Actor, int64, error) {
	ctx := context.Background()
	key := fmt.Sprintf("actorsList:size=%v:currPage=%v", perPage, currentPage)
	cache, err := global.Redis.HGetAll(ctx, key).Result()
	if err == nil {
		var cacheActors []models.Actor
		err = json.Unmarshal([]byte(cache["actors"]), &cacheActors)
		if err != nil {
			global.Logger.Warn("json unmarshal error", err)
		} else {
			count, _ := strconv.ParseInt(cache["count"], 10, 64)
			return cacheActors, count, nil
		}
	}

	actors, count, actorsErr := service.repo.Index(currentPage, perPage)
	if actorsErr == nil {
		bytes, err := json.Marshal(actors)
		if err != nil {
			global.Logger.Warn("json marshal error", err)
		} else {
			global.Redis.HSet(ctx, key, "actors", bytes, "count", count)
			global.Redis.Expire(ctx, key, 24*time.Hour)
		}
	}
	return actors, count, actorsErr
}

func (service *ActorService) Show(id string) (models.Actor, error) {
	ctx := context.Background()
	key := fmt.Sprintf("actor:%v", id)
	bytes, err := global.Redis.Get(ctx, key).Bytes()
	if err == nil {
		var cacheActor models.Actor
		err = json.Unmarshal(bytes, &cacheActor)
		if err != nil {
			global.Logger.Warn("json unmarshal error", err)
		} else {
			return cacheActor, nil
		}
	}

	actor, actorErr := service.repo.Show(id)
	if actorErr == nil {
		bytes, err = json.Marshal(actor)
		if err != nil {
			global.Logger.Warn("json marshal error", err)
		} else {
			global.Redis.Set(ctx, key, bytes, 24*time.Hour)
		}
	}
	return actor, actorErr
}

func (service *ActorService) Store(request *request.Actor) (models.Actor, error) {
	actor := models.Actor{
		Name:           request.Name,
		TranslatedName: request.TranslatedName,
		NickName:       request.NickName,
		Nationality:    request.Nationality,
		Born:           request.Born,
		ImageId:        request.ImageId,
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
		ImageId:        request.ImageId,
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
	ctx := context.Background()
	key := fmt.Sprintf("actor:%v", id)
	global.Redis.Del(ctx, key)
}

func (service *ActorService) delActorsListCache() {
	ctx := context.Background()
	iter := global.Redis.Scan(ctx, 0, "actorsList:size=*:currPage=*", 0).Iterator()
	for iter.Next(ctx) {
		global.Redis.Del(ctx, iter.Val())
	}
}
