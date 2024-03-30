package initialize

import (
	"musical_wiki/config"
	"sync"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var dbOnce sync.Once
var dbErr error

func NewDB(config config.Pg, logger *zap.SugaredLogger) *gorm.DB {
	if db == nil {
		dbOnce.Do(func() {
			db, dbErr = gorm.Open(postgres.Open(config.Dsn()), &gorm.Config{})
			if dbErr != nil || db == nil {
				logger.Error("db init error ", dbErr)
			}
		})
	}
	return db
}
