package repository

import (
	"musical_wiki/global"
	"musical_wiki/models"
)

type ImageRepository struct{}

func (repository *ImageRepository) IndexGallery(actorId string) ([]models.Image, error) {
	var images []models.Image
	err := global.Db.Where("actor_id = ?", actorId).Where("image_type = ?", "GALLERY").Find(&images).Error
	if err != nil {
		return nil, err
	}
	return images, err
}

func (repository *ImageRepository) ShowAvatar(actorId string) (models.Image, error) {
	var image = models.Image{}
	err := global.Db.Where("actor_id = ?", actorId).Where("image_type = ?", "AVATAR").First(&image).Error
	return image, err
}

func (repository *ImageRepository) Show(id string) (models.Image, error) {
	var image = models.Image{}
	err := global.Db.First(&image, id).Error
	return image, err
}

func (repository *ImageRepository) Store(image *models.Image) error {
	// Check if actor exists
	var actor models.Actor
	err := global.Db.Where("id = ?", image.ActorId).First(&actor).Error
	if err != nil {
		return err
	}
	return global.Db.Create(&image).Error
}

func (repository *ImageRepository) Update(id string, image *models.Image) error {
	var model models.Image
	err := global.Db.Where("id = ?", id).First(&model).Error
	if err != nil {
		return err
	}
	err = global.Db.Model(&model).Updates(image).Error
	return err
}

func (repository *ImageRepository) Destroy(id string) error {
	return global.Db.Delete(&models.Image{}, id).Error
}
