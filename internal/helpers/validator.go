package helpers

import (
	"log"
	"sync"

	"github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	es_translations "github.com/go-playground/validator/v10/translations/es"
)

var (
	once       sync.Once
	Validate   *validator.Validate
	Translator ut.Translator
)

func InitValidator() {
	once.Do(func() {
		Validate = validator.New()

		spanish := es.New()
		uni := ut.New(spanish, spanish)

		var found bool
		Translator, found = uni.GetTranslator("es")
		if !found {
			log.Panic("Not found spanish translator")
		}

		err := es_translations.RegisterDefaultTranslations(Validate, Translator)
		if err != nil {
			log.Panic("Cannot set spanish Transalator ", err)
		}
	})
}

func ValidatorErrorsMap(err error) []string {
	var errorFields []string

	for _, err := range err.(validator.ValidationErrors) {
		errorFields = append(errorFields, err.Translate(Translator))
	}

	return errorFields
}
