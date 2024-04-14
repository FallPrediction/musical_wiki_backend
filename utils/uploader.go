package utils

import (
	"bytes"
	"context"
	"fmt"
	"musical_wiki/config"
	"musical_wiki/helper"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.uber.org/zap"
)

type Uploader struct {
	s3     *s3.Client
	logger *zap.SugaredLogger
}

func (u *Uploader) Upload(file *helper.File) (string, error) {
	uploader := manager.NewUploader(u.s3)
	imageDecode, err := file.Decode()
	if err != nil {
		u.logger.Error("File decode error", err)
		return "", err
	}
	fileName := fmt.Sprint(helper.NewStr().Random(10), file.GetExt())
	output, err := uploader.Upload(context.Background(), &s3.PutObjectInput{
		Bucket: &config.Bucket,
		Key:    &fileName,
		Body:   bytes.NewReader(imageDecode),
	})
	if err != nil {
		u.logger.Error("File upload error", err)
		return "", err
	}
	return *output.Key, err
}

func (u *Uploader) Delete(key string) error {
	_, err := u.s3.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: &config.Bucket,
		Key:    &key,
	})
	if err != nil {
		u.logger.Error("File delete error", err)
	}
	return err
}

func NewUploader(s3 *s3.Client, logger *zap.SugaredLogger) Uploader {
	return Uploader{s3: s3, logger: logger}
}
