package utils

import (
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

func PrepareValidator(validator *validator.Validate) *validator.Validate {
	validator = ValidatorNewWithTagNameInJson(validator)
	validator = ValidatorRegisterValidationRFC3339Nano(validator)
	validator = ValidatorRegisterValidationRegexp(validator)

	return validator
}

func ValidatorRFC3339NanoFn(fl validator.FieldLevel) bool {
	val := fl.Field().String()

	if val == "" {
		return true
	}

	if _, err := time.Parse(time.RFC3339Nano, val); err == nil {
		return true
	}

	return false
}

func ValidatorRegisterValidationRFC3339Nano(validator *validator.Validate) *validator.Validate {
	if err := validator.RegisterValidation("rfc3339nano", ValidatorRFC3339NanoFn, false); err != nil {
		panic(err)
	}

	return validator
}

func ValidatorRegexpFn(fl validator.FieldLevel) bool {
	match, _ := regexp.MatchString(`^\d+-\d+$`, fl.Field().String())
	return match
}

func ValidatorRegisterValidationRegexp(validator *validator.Validate) *validator.Validate {
	if err := validator.RegisterValidation("regexp", ValidatorRegexpFn, false); err != nil {
		panic(err)
	}

	return validator
}

func ValidatorNewWithTagNameInJson(validator *validator.Validate) *validator.Validate {
	validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	return validator
}
