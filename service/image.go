package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"musical_wiki/config"
	"musical_wiki/global"
	"musical_wiki/helper"
	"musical_wiki/models"
	"musical_wiki/repository"
	"musical_wiki/request"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"gorm.io/gorm"
)

type ImageService struct {
	repo repository.ImageRepository
}

func (service *ImageService) IndexGallery(actorId string) ([]models.Image, error) {
	key := fmt.Sprint("imageGallery:actorId=", actorId)
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	cacheBytes, err := global.Redis.Get(ctx, key).Bytes()
	if err == nil {
		var cacheImages []models.Image
		err = json.Unmarshal(cacheBytes, &cacheImages)
		if err != nil {
			global.Logger.Warn("json unmarshal error", err)
		} else {
			return cacheImages, nil
		}
	}
	cancel()

	images, imagesErr := service.repo.IndexGallery(actorId)
	if imagesErr == nil {
		cacheBytes, err = json.Marshal(images)
		if err != nil {
			global.Logger.Warn("json marshal error", err)
		} else {
			ctx, cancel = context.WithTimeout(context.Background(), 500*time.Millisecond)
			defer cancel()
			global.Redis.Set(ctx, key, cacheBytes, 24*time.Hour)
		}
	}
	return images, imagesErr
}

func (service *ImageService) ShowAvatar(actorId string) (models.Image, error) {
	key := fmt.Sprint("imageAvatar:actorId=", actorId)
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	cacheBytes, err := global.Redis.Get(ctx, key).Bytes()
	if err == nil {
		var cacheImage models.Image
		err = json.Unmarshal(cacheBytes, &cacheImage)
		if err != nil {
			global.Logger.Warn("json unmarshal error", err)
		} else {
			return cacheImage, nil
		}
	}
	cancel()

	image, imagesErr := service.repo.ShowAvatar(actorId)
	if imagesErr == nil {
		cacheBytes, err = json.Marshal(image)
		if err != nil {
			global.Logger.Warn("json marshal error", err)
		} else {
			ctx, cancel = context.WithTimeout(context.Background(), 500*time.Millisecond)
			defer cancel()
			global.Redis.Set(ctx, key, cacheBytes, 24*time.Hour)
		}
	}
	return image, imagesErr
}

func (service *ImageService) UpdateAvatar(request *request.Image) (models.Image, error) {
	// Upload new avatar
	uploader := manager.NewUploader(global.S3)
	file := helper.NewFile(request.Name, request.Image)
	imageDecode, err := file.Decode()
	if err != nil {
		return models.Image{}, err
	}
	fileName := fmt.Sprint(helper.NewStr().Random(10), file.GetExt())
	_, err = uploader.Upload(context.Background(), &s3.PutObjectInput{
		Bucket: &config.Bucket,
		Key:    &fileName,
		Body:   bytes.NewReader(imageDecode),
	})
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
			_, deleteErr := global.S3.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
				Bucket: &config.Bucket,
				Key:    &oldImage.ImageName,
			})
			if deleteErr != nil {
				global.Logger.Warn("S3 remove failed: ", oldImage.ImageName)
			}
		}
	}

	service.delImageCache(fmt.Sprint(request.ActorId))
	actorService := ActorService{}
	actorService.delActorCache(fmt.Sprint(request.ActorId))
	actorService.delActorsListCache()
	return image, imageErr
}

func (service *ImageService) StoreGallery(request *request.Image) (models.Image, error) {
	// Upload new image
	uploader := manager.NewUploader(global.S3)
	file := helper.NewFile(request.Name, request.Image)
	imageDecode, err := file.Decode()
	if err != nil {
		return models.Image{}, err
	}
	fileName := fmt.Sprint(helper.NewStr().Random(10), file.GetExt())
	_, err = uploader.Upload(context.Background(), &s3.PutObjectInput{
		Bucket: &config.Bucket,
		Key:    &fileName,
		Body:   bytes.NewReader(imageDecode),
	})
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
	actorService := ActorService{}
	actorService.delActorCache(fmt.Sprint(request.ActorId))
	actorService.delActorsListCache()
	return image, imageErr
}

func (service *ImageService) Destroy(id string) error {
	image, err := service.repo.Show(id)
	if err != nil {
		return err
	}

	global.S3.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: &config.Bucket,
		Key:    &image.ImageName,
	})

	service.delImageCache(fmt.Sprint(image.ActorId))
	actorService := ActorService{}
	actorService.delActorCache(fmt.Sprint(image.ActorId))
	actorService.delActorsListCache()
	return service.repo.Destroy(id)
}

func (service *ImageService) delImageCache(actorId string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	key := fmt.Sprintf("imageAvatar:actorId=%v", actorId)
	global.Redis.Del(ctx, key)
	if ctx.Err() == context.DeadlineExceeded {
		global.Logger.Warn("delAvatarCache timeout", key)
	}
	cancel()
	key = fmt.Sprintf("imageGallery:actorId=%v", actorId)
	global.Redis.Del(ctx, key)
	if ctx.Err() == context.DeadlineExceeded {
		global.Logger.Warn("delGalleryCache timeout", key)
	}
	cancel()
}
