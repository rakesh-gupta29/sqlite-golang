package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func isName(fl validator.FieldLevel) bool {
	err := validate.Var(fl, "required,alphanum")
	if err != nil {
		return false
	}
	return true
}

// mounting custom validator logics
func MountValidators(app *fiber.App) {
	validate.RegisterValidation("is-name", isName)
}

func String(str string) (bool, string) {
	// trimmedValue := strings.TrimSpace(str)

	err := validate.Var(str, "required")
	if err != nil {
		return false, "Invalid string"
	}
	return true, ""
}

func Required(value, fieldName string) (bool, string) {
	err := validate.Var(value, "required")
	if err != nil {
		return false, "fielasdfd is required"
	}
	return true, ""
}

func Email(str string) (bool, string) {
	err := validate.Var(str, "required,email")
	if err != nil {
		return false, "Email is invalid"
	}
	return true, ""
}
