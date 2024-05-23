package validations

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

// Проверяет кейс: Если текущее поле содержит контент, то проверяется зависимое поле, если оба поля пустые, то игнорируются
func RequireAnotherField(fl validator.FieldLevel) bool {
	paramField := fl.Param()

	if paramField == `` {
		return true
	}

	var paramFieldValue reflect.Value

	if fl.Parent().Kind() == reflect.Ptr {
		paramFieldValue = fl.Parent().Elem().FieldByName(paramField)
	} else {
		paramFieldValue = fl.Parent().FieldByName(paramField)
	}

	selfValue := fl.Field().String()

	value := paramFieldValue.String()

	if selfValue == "" && value == "" {
		return true
	}

	return value != ""
}
