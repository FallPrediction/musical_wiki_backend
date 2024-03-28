package initialize

import (
	"musical_wiki/config"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	err error
)

func NewDB(config config.Pg, logger *zap.SugaredLogger) *gorm.DB {
	db, _ := gorm.Open(postgres.Open(config.Dsn()), &gorm.Config{})
	if err != nil || db == nil {
		logger.Error("db init error ", err)
	}
	return db
}
