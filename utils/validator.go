package utils

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// 定义一个全局翻译器T
var Trans ut.Translator

var (
	uni      *ut.UniversalTranslator
	Validate *validator.Validate
)

func InputValidate[T any](obj interface{}, res *Result[T]) error {
	err = Validate.Struct(obj)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, v := range errs {
			s := v.Translate(Trans)
			res.Code = 500
			res.Message = s
			res.IsSuccess = false
			return v
		}

	}
	return nil
}

// InitTrans 初始化翻译器
func InitTrans(locale string) (trans ut.Translator, err error) {
	en := en.New()
	uni = ut.New(en, en)
	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	Trans, _ = uni.GetTranslator(locale)
	Validate = validator.New()
	err = en_translations.RegisterDefaultTranslations(Validate, Trans)

	return Trans, err
}
