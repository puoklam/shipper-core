package validation

import "github.com/go-playground/validator/v10"

var Vali *validator.Validate

func init() {
	Vali = validator.New()
}

type Validator[T any] interface {
	From(T)
	Done()
}
