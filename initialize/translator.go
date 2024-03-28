package initialize

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTwTrans "github.com/go-playground/validator/v10/translations/zh_tw"
	"go.uber.org/zap"
)

func NewTranslator(logger *zap.SugaredLogger) ut.Translator {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 翻譯器
		enT := en.New()
		twT := zh_Hant_TW.New()

		// 英文fallback，繁體中文支援
		universalTranslator := ut.New(enT, twT)
		translator, ok := universalTranslator.GetTranslator("zh_Hant_TW")
		if !ok {
			logger.Error("Get universal translator zh_Hant_TW failed")
		}

		// 註冊翻譯器到驗證器
		zhTwTrans.RegisterDefaultTranslations(v, translator)
		return translator
	}
	return nil
}
