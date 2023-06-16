package util

import (
	"github.com/go-playground/validator/v10"
	"inventory-management/backend/internal/http/presenter/response"
	"reflect"
)

var validate = validator.New()

func ValidateStruct(entity interface{}) []*response.ErrorResponse {
	val := reflect.ValueOf(entity)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return []*response.ErrorResponse{}
	}

	err := validate.Struct(entity)
	if err != nil {
		var errors []*response.ErrorResponse
		for _, err := range err.(validator.ValidationErrors) {
			var element response.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}

		return errors
	}

	return nil
}
