package model

import validation "github.com/go-ozzo/ozzo-validation"

// Вызывается только в случае переданного true условия
func requiredIf(condition bool) validation.RuleFunc {
	return func(value interface{}) error {
		if condition {
			// Напрямую достаем из библиотеки Validate c правилом required
			return validation.Validate(value, validation.Required)
		}

		return nil
	}
}
