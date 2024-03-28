package initialize

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.uber.org/zap"
)

func NewS3(logger *zap.SugaredLogger) *s3.Client {
	cfg, err := awsConfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		logger.Error("unable to load SDK config", err)
	}
	return s3.NewFromConfig(cfg)
}
