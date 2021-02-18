package validator

import (
	"reflect"

	ut "github.com/go-playground/universal-translator"
	playgroundvalidator "github.com/go-playground/validator/v10"
	translationszh "github.com/go-playground/validator/v10/translations/zh"
)

type validator struct {
}

func NewValidator() *validator {
	return &validator{}
}

func (vd *validator) Register(v *playgroundvalidator.Validate, trans ut.Translator) error {

	_ = v.RegisterValidation("test", func(fl playgroundvalidator.FieldLevel) bool {
		str, ok := fl.Field().Interface().(string)
		if ok {
			if str != "test" {
				return false
			}
		}
		return true
	})

	_ = translationszh.RegisterDefaultTranslations(v, trans)

	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			return field.Name
		}
		return label
	})

	_ = v.RegisterTranslation("test", trans, func(ut ut.Translator) error {
		return ut.Add("test", "{0}测试检查", true)
	}, func(ut ut.Translator, fe playgroundvalidator.FieldError) string {
		t, _ := ut.T("test", fe.Field())
		return t
	})

	return nil
}
