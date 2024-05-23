package validations

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Проверяет дату\время в формате "YYYY-MM-DDThh:mm:ss.fffZ"
func ValidateISODate(fl validator.FieldLevel) bool {
	ISO8601DateRegexString := "^(\\d{4})(-(0[1-9]|1[0-2])(-([12]\\d|0[1-9]|3[01]))([T\\s]((([01]\\d|2[0-3])((:)[0-5]\\d))([\\:]\\d+)?)?(:[0-5]\\d([\\.]\\d+)?)?([zZ]|([\\+-])([01]\\d|2[0-3]):?([0-5]\\d)?)?)?)$"
	ISO8601DateRegex := regexp.MustCompile(ISO8601DateRegexString)
	return ISO8601DateRegex.MatchString(fl.Field().String())
}
