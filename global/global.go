package global

import (
	ut "github.com/go-playground/universal-translator"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Db         *gorm.DB
	Logger     *zap.SugaredLogger
	Translator ut.Translator
)
