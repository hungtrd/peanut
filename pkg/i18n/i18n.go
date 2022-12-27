package i18n

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/ja"
	"github.com/go-playground/locales/vi"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enVt "github.com/go-playground/validator/v10/translations/en"
	jaVt "github.com/go-playground/validator/v10/translations/ja"
	viVt "github.com/go-playground/validator/v10/translations/vi"
)

var uni *ut.UniversalTranslator

func SetupI18n() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// Default language order fisrt when init
		uni = ut.New(en.New(), vi.New(), ja.New())

		trans, _ := uni.GetTranslator("vi")
		_ = viVt.RegisterDefaultTranslations(v, trans)

		trans, _ = uni.GetTranslator("en")
		_ = enVt.RegisterDefaultTranslations(v, trans)

		trans, _ = uni.GetTranslator("ja")
		_ = jaVt.RegisterDefaultTranslations(v, trans)
	}
}

func GetTrans(locale string) ut.Translator {
	trans, _ := uni.GetTranslator(locale)
	return trans
}
