package api

import (
	"github.com/BMogetta/goBank_go-course/util"
	"github.com/go-playground/validator/v10"
)

// check if curency is supported
var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {

	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
