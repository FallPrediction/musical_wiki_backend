package service

import (
	"errors"
	"fmt"
	"musical_wiki/helper"
	"musical_wiki/models"
	"musical_wiki/repository"
	"musical_wiki/request"
	"musical_wiki/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ImageService struct {
	repo     repository.ImageRepository
	logger   *zap.SugaredLogger
	uploader utils.Uploader
	cache    utils.Cache
}

func (service *ImageService) IndexGallery(actorId string) ([]models.Image, error) {
	key := fmt.Sprint("imageGallery:actorId=", actorId)
	var cacheImages []models.Image
	cacheErr := service.cache.Get(key, &cacheImages)
	if cacheErr == nil {
		return cacheImages, nil
	}

	images, imagesErr := service.repo.IndexGallery(actorId)
	if imagesErr == nil {
		service.cache.Set(key, images)
	}
	return images, imagesErr
}

func (service *ImageService) ShowAvatar(actorId string) (models.Image, error) {
	key := fmt.Sprint("imageAvatar:actorId=", actorId)
	var cacheImage models.Image
	cacheErr := service.cache.Get(key, &cacheImage)
	if cacheErr == nil {
		return cacheImage, nil
	}

	image, imagesErr := service.repo.ShowAvatar(actorId)
	if imagesErr == nil {
		service.cache.Set(key, image)
	}
	return image, imagesErr
}

func (service *ImageService) UpdateAvatar(request *request.Image) (models.Image, error) {
	// Upload new avatar
	file := helper.NewFile(request.Name, request.Image)
	fileName, err := service.uploader.Upload(file)
	if err != nil {
		return models.Image{}, err
	}

	// Update or create image record
	image := models.Image{
		ImageName: fileName,
		Mime:      file.GetMime(),
		ActorId:   request.ActorId,
		ImageType: "AVATAR",
	}
	oldImage, oldImageErr := service.repo.ShowAvatar(fmt.Sprint(request.ActorId))
	var imageErr error
	if errors.Is(oldImageErr, gorm.ErrRecordNotFound) {
		imageErr = service.repo.Store(&image)
	} else {
		imageErr = service.repo.Update(fmt.Sprint(oldImage.Id), &image)

		// Delete old avatar
		if oldImageErr == nil {
			deleteErr := service.uploader.Delete(oldImage.ImageName)
			if deleteErr != nil {
				service.logger.Warn("S3 remove failed: ", oldImage.ImageName)
			}
		}
	}

	service.delImageCache(fmt.Sprint(request.ActorId))
	service.delActorCache(fmt.Sprint(request.ActorId))
	service.delActorsListCache()
	return image, imageErr
}

func (service *ImageService) StoreGallery(request *request.Image) (models.Image, error) {
	// Upload new image
	file := helper.NewFile(request.Name, request.Image)
	fileName, err := service.uploader.Upload(file)
	if err != nil {
		return models.Image{}, err
	}

	// Create image record
	image := models.Image{
		ImageName: fileName,
		Mime:      file.GetMime(),
		ActorId:   request.ActorId,
		ImageType: "GALLERY",
	}
	imageErr := service.repo.Store(&image)

	service.delImageCache(fmt.Sprint(request.ActorId))
	service.delActorCache(fmt.Sprint(request.ActorId))
	service.delActorsListCache()
	return image, imageErr
}

func (service *ImageService) Destroy(id string) error {
	image, err := service.repo.Show(id)
	if err != nil {
		return err
	}

	deleteErr := service.uploader.Delete(image.ImageName)
	if deleteErr != nil {
		service.logger.Warn("S3 remove failed: ", image.ImageName)
	}

	service.delImageCache(fmt.Sprint(image.ActorId))
	service.delActorCache(fmt.Sprint(image.ActorId))
	service.delActorsListCache()
	return service.repo.Destroy(id)
}

func (service *ImageService) delImageCache(actorId string) {
	service.cache.Del(fmt.Sprintf("imageAvatar:actorId=%v", actorId))
	service.cache.Del(fmt.Sprintf("imageGallery:actorId=%v", actorId))
}

func (service *ImageService) delActorCache(id string) {
	service.cache.Del(fmt.Sprintf("actor:%v", id))
}

func (service *ImageService) delActorsListCache() {
	service.cache.ScanAndDel("actorsList:size=*:currPage=*")
}

func NewImageService(repo repository.ImageRepository, logger *zap.SugaredLogger, uploader utils.Uploader, cache utils.Cache) ImageService {
	return ImageService{repo: repo, logger: logger, uploader: uploader, cache: cache}
}
