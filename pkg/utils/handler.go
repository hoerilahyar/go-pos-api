package utils

import (
	"errors"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func HandleValidationError(err error, obj any) []string {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		var errorMessages []string
		objType := reflect.TypeOf(obj)
		if objType.Kind() == reflect.Ptr {
			objType = objType.Elem()
		}

		for _, e := range ve {
			// Dapatkan nama field JSON dari struct
			if f, ok := objType.FieldByName(e.StructField()); ok {
				jsonTag := f.Tag.Get("json")
				if jsonTag == "" || jsonTag == "-" {
					jsonTag = e.Field() // fallback
				}
				errorMessages = append(errorMessages, jsonTag+" is required")
			} else {
				errorMessages = append(errorMessages, e.Field()+" is required")
			}
		}
		return errorMessages
	}
	return nil
}
