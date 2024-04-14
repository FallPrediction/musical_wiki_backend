package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"musical_wiki/helper"
	"musical_wiki/models"
	"musical_wiki/repository"
	"musical_wiki/request"
	"musical_wiki/utils"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ImageService struct {
	repo     repository.ImageRepository
	logger   *zap.SugaredLogger
	redis    *redis.Client
	uploader utils.Uploader
}

func (service *ImageService) IndexGallery(actorId string) ([]models.Image, error) {
	key := fmt.Sprint("imageGallery:actorId=", actorId)
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	cacheBytes, err := service.redis.Get(ctx, key).Bytes()
	if err == nil {
		var cacheImages []models.Image
		err = json.Unmarshal(cacheBytes, &cacheImages)
		if err != nil {
			service.logger.Warn("json unmarshal error", err)
		} else {
			return cacheImages, nil
		}
	}
	cancel()

	images, imagesErr := service.repo.IndexGallery(actorId)
	if imagesErr == nil {
		cacheBytes, err = json.Marshal(images)
		if err != nil {
			service.logger.Warn("json marshal error", err)
		} else {
			ctx, cancel = context.WithTimeout(context.Background(), 500*time.Millisecond)
			defer cancel()
			service.redis.Set(ctx, key, cacheBytes, 24*time.Hour)
		}
	}
	return images, imagesErr
}

func (service *ImageService) ShowAvatar(actorId string) (models.Image, error) {
	key := fmt.Sprint("imageAvatar:actorId=", actorId)
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	cacheBytes, err := service.redis.Get(ctx, key).Bytes()
	if err == nil {
		var cacheImage models.Image
		err = json.Unmarshal(cacheBytes, &cacheImage)
		if err != nil {
			service.logger.Warn("json unmarshal error", err)
		} else {
			return cacheImage, nil
		}
	}
	cancel()

	image, imagesErr := service.repo.ShowAvatar(actorId)
	if imagesErr == nil {
		cacheBytes, err = json.Marshal(image)
		if err != nil {
			service.logger.Warn("json marshal error", err)
		} else {
			ctx, cancel = context.WithTimeout(context.Background(), 500*time.Millisecond)
			defer cancel()
			service.redis.Set(ctx, key, cacheBytes, 24*time.Hour)
		}
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	key := fmt.Sprintf("imageAvatar:actorId=%v", actorId)
	service.redis.Del(ctx, key)
	if ctx.Err() == context.DeadlineExceeded {
		service.logger.Warn("delAvatarCache timeout", key)
	}
	cancel()
	key = fmt.Sprintf("imageGallery:actorId=%v", actorId)
	service.redis.Del(ctx, key)
	if ctx.Err() == context.DeadlineExceeded {
		service.logger.Warn("delGalleryCache timeout", key)
	}
	cancel()
}

func (service *ImageService) delActorCache(id string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	key := fmt.Sprintf("actor:%v", id)
	service.redis.Del(ctx, key)
	if ctx.Err() == context.DeadlineExceeded {
		service.logger.Warn("delActorCache timeout", key)
	}
}

func (service *ImageService) delActorsListCache() {
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

func NewImageService(repo repository.ImageRepository, logger *zap.SugaredLogger, redis *redis.Client, uploader utils.Uploader) ImageService {
	return ImageService{repo: repo, logger: logger, redis: redis, uploader: uploader}
}
