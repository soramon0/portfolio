package lib

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/soramon0/portfolio/src/internal/types"
)

type ValidatorTranslator struct {
	Validator    *validator.Validate
	Translations ut.Translator
}

func NewValidator() (*ValidatorTranslator, error) {
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	validate := validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)

	if err := registerOverrides(trans, validate); err != nil {
		return nil, err
	}

	return &ValidatorTranslator{
		Validator:    validate,
		Translations: trans,
	}, nil
}

func (vt *ValidatorTranslator) ValidationErrors(ve validator.ValidationErrors) *types.APIValidationErrors {
	out := make([]types.APIFieldError, len(ve))
	for i, fe := range ve {
		t := ve.Translate(vt.Translations)
		out[i] = types.APIFieldError{Field: fe.Field(), Message: t[fe.Namespace()]}
	}

	return &types.APIValidationErrors{Errors: out}
}

func registerOverrides(trans ut.Translator, v *validator.Validate) error {
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is required", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})
}
