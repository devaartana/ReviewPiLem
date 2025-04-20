package utils

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var validate = validator.New()

func init() {
	validate.RegisterValidation("username", func(fl validator.FieldLevel) bool {
		username := fl.Field().String()
		re := regexp.MustCompile(`^[a-z0-9_-]{3,20}$`)
		return re.MatchString(username)
	})
}

func ValidateUsername(username string) error{
	return validate.Var(username, "required,username")
}

func ValidateEmail(email string) error{
	return validate.Var(email, "required,email")
}
