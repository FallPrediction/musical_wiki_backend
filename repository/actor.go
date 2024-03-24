package repository

import (
	"musical_wiki/global"
	"musical_wiki/models"
)

type ActorRepository struct{}

func (repository *ActorRepository) Index(currentPage int, perPage int) ([]models.Actor, int64, error) {
	var actors []models.Actor
	var count int64
	global.Db.Model(&models.Actor{}).Count(&count)
	err := global.Db.Order("id").Limit(perPage).Offset((currentPage-1)*perPage).Preload("Avatar", "image_type = ?", "AVATAR").Find(&actors).Error
	return actors, count, err
}

func (repository *ActorRepository) Show(id string) (models.Actor, error) {
	var actor models.Actor
	err := global.Db.Where("id = ?", id).Preload("Avatar", "image_type = ?", "AVATAR").First(&actor).Error
	return actor, err
}

func (repository *ActorRepository) Store(actor *models.Actor) error {
	return global.Db.Create(&actor).Error
}

func (repository *ActorRepository) Update(id string, actor *models.Actor) error {
	var model *models.Actor
	err := global.Db.Where("id = ?", id).First(&model).Error
	if err != nil {
		return err
	}
	err = global.Db.Model(&model).Updates(actor).Error
	return err
}

func (repository *ActorRepository) Destroy(id string) error {
	var actor models.Actor
	err := global.Db.Where("id = ?", id).First(&actor).Error
	if err != nil {
		return err
	}
	err = global.Db.Delete(&actor).Error
	return err
}
