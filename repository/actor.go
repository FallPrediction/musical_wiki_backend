package repository

import (
	"musical_wiki/models"

	"gorm.io/gorm"
)

type ActorRepository struct {
	db *gorm.DB
}

func (repository *ActorRepository) Index(currentPage int, perPage int) ([]models.Actor, int64, error) {
	var actors []models.Actor
	var count int64
	repository.db.Model(&models.Actor{}).Count(&count)
	err := repository.db.Order("id").Limit(perPage).Offset((currentPage-1)*perPage).Preload("Avatar", "image_type = ?", "AVATAR").Find(&actors).Error
	return actors, count, err
}

func (repository *ActorRepository) Show(id string) (models.Actor, error) {
	var actor models.Actor
	err := repository.db.Where("id = ?", id).Preload("Avatar", "image_type = ?", "AVATAR").First(&actor).Error
	return actor, err
}

func (repository *ActorRepository) Store(actor *models.Actor) error {
	return repository.db.Create(&actor).Error
}

func (repository *ActorRepository) Update(id string, actor *models.Actor) error {
	var model *models.Actor
	err := repository.db.Where("id = ?", id).First(&model).Error
	if err != nil {
		return err
	}
	err = repository.db.Model(&model).Updates(actor).Error
	return err
}

func (repository *ActorRepository) Destroy(id string) error {
	var actor models.Actor
	err := repository.db.Where("id = ?", id).First(&actor).Error
	if err != nil {
		return err
	}
	err = repository.db.Delete(&actor).Error
	return err
}

func NewActorRepository(db *gorm.DB) ActorRepository {
	return ActorRepository{db}
}
