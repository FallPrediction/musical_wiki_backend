package global

import (
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
)
