package initialize

import (
	"musical_wiki/config"
	"musical_wiki/global"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	err error
)

func InitDatabase() {
	global.Db, err = gorm.Open(postgres.Open(config.NewPg().Dsn()), &gorm.Config{})
	if err != nil || global.Db == nil {
		global.Logger.Error("db init error", err)
	}
}
