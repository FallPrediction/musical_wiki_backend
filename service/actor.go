package service

import (
	"musical_wiki/models"
	"musical_wiki/repository"
	"musical_wiki/request"
)

type ActorService struct {
	repo repository.ActorRepository
}

func (service *ActorService) Index() ([]models.Actor, error) {
	return service.repo.Index()
}

func (service *ActorService) Show(id string) (models.Actor, error) {
	return service.repo.Show(id)
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
	return service.repo.Update(id, &actor)
}

func (service *ActorService) Destroy(id string) error {
	return service.repo.Destroy(id)
}
