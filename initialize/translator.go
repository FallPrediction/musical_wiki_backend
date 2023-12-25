package initialize

import (
	"musical_wiki/global"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTwTrans "github.com/go-playground/validator/v10/translations/zh_tw"
)

func InitTranslator() error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 翻譯器
		enT := en.New()
		twT := zh_Hant_TW.New()

		// 英文fallback，繁體中文支援
		universalTranslator := ut.New(enT, twT)
		global.Translator, ok = universalTranslator.GetTranslator("zh_Hant_TW")
		if !ok {
			global.Logger.Error("Get universal translator zh_Hant_TW failed")
		}

		// 註冊翻譯器到驗證器
		return zhTwTrans.RegisterDefaultTranslations(v, global.Translator)
	}
	return nil
}
