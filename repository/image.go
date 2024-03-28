package repository

import (
	"musical_wiki/models"

	"gorm.io/gorm"
)

type ImageRepository struct {
	db *gorm.DB
}

func (repository *ImageRepository) IndexGallery(actorId string) ([]models.Image, error) {
	var images []models.Image
	err := repository.db.Where("actor_id = ?", actorId).Where("image_type = ?", "GALLERY").Find(&images).Error
	if err != nil {
		return nil, err
	}
	return images, err
}

func (repository *ImageRepository) ShowAvatar(actorId string) (models.Image, error) {
	var image = models.Image{}
	err := repository.db.Where("actor_id = ?", actorId).Where("image_type = ?", "AVATAR").First(&image).Error
	return image, err
}

func (repository *ImageRepository) Show(id string) (models.Image, error) {
	var image = models.Image{}
	err := repository.db.First(&image, id).Error
	return image, err
}

func (repository *ImageRepository) Store(image *models.Image) error {
	// Check if actor exists
	var actor models.Actor
	err := repository.db.Where("id = ?", image.ActorId).First(&actor).Error
	if err != nil {
		return err
	}
	return repository.db.Create(&image).Error
}

func (repository *ImageRepository) Update(id string, image *models.Image) error {
	var model models.Image
	err := repository.db.Where("id = ?", id).First(&model).Error
	if err != nil {
		return err
	}
	err = repository.db.Model(&model).Updates(image).Error
	return err
}

func (repository *ImageRepository) Destroy(id string) error {
	return repository.db.Delete(&models.Image{}, id).Error
}

func NewImageRepository(db *gorm.DB) ImageRepository {
	return ImageRepository{db: db}
}
