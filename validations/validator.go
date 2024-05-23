package validations

import (
	"errors"
	"fmt"

	"github.com/go-playground/locales/ru"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ru_translations "github.com/go-playground/validator/v10/translations/ru"
)

/*
Корневой валидатор, содержащий все основные и дополнительные валидации.
Также содержит в себе настроенную локализацию на русский язык.
*/
type CoreValidator struct {
	Validator   *validator.Validate
	Translation ut.Translator
}

// Создает новый экземпляр корневого валидатора
func NewValidator() (v *CoreValidator, err error) {
	v = &CoreValidator{
		Validator: validator.New(),
	}

	ru := ru.New()
	uni := ut.New(ru, ru)
	trans, _ := uni.GetTranslator("ru")

	v.Translation = trans

	err = ru_translations.RegisterDefaultTranslations(v.Validator, trans)

	if err != nil {
		return
	}

	err = v.Validator.RegisterValidation("isoTime", ValidateISODate)

	if err != nil {
		return
	}

	err = v.Validator.RegisterValidation("enum", ValidateEnum)

	if err != nil {
		return
	}

	err = v.Validator.RegisterValidation("req", RequireAnotherField)

	if err != nil {
		return
	}

	err = v.Validator.RegisterTranslation(
		"isoTime",
		v.Translation,
		func(ut ut.Translator) error {
			return ut.Add("isoTime", "{0} содержит некорретный формат даты", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("isoTime", fe.Field())
			return t
		},
	)

	return
}

// Валидирует структуру, в случае ошибок возвращает локализированную строку
func (cv *CoreValidator) Validate(i interface{}) (err error) {
	err = cv.Validator.Struct(i)
	if err != nil {
		msg := ""
		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			msg = fmt.Sprint(msg, " ", e.Translate(cv.Translation), ";")
		}
		err = errors.New(msg)
	}

	return
}
