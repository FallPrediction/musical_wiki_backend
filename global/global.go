package global

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	ut "github.com/go-playground/universal-translator"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Db         *gorm.DB
	Logger     *zap.SugaredLogger
	Translator ut.Translator
	Redis      *redis.Client
	S3         *s3.Client
)
