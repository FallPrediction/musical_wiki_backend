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

type ActorService struct {
	repo repository.ActorRepository
}

func (service *ActorService) Index() ([]models.Actor, error) {
	return service.repo.Index()
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

	ctx := context.Background()
	key := fmt.Sprintf("actor:%v", id)
	global.Redis.Del(ctx, key)

	return service.repo.Update(id, &actor)
}

func (service *ActorService) Destroy(id string) error {
	return service.repo.Destroy(id)
}
