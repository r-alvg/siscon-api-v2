package helpers

import "github.com/go-playground/validator/v10"

type err struct {
	Key   string
	Cause string
}

func ValidateErros(errs []validator.FieldError) []err {
	var ers []err
	for _, v := range errs {
		ers = append(ers, err{v.Field(), v.Tag()})
	}
	return ers
}

