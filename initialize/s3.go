package initialize

import (
	"context"
	"musical_wiki/global"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func InitS3() {
	cfg, err := awsConfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		global.Logger.Error("unable to load SDK config", err)
	}
	global.S3 = s3.NewFromConfig(cfg)
}
