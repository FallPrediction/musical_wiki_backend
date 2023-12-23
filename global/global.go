package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Db     *gorm.DB
	Logger *zap.SugaredLogger
)
