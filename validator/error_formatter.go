package validator

import (
	"encoding/json"
	"errors"
	"io"
	"strings"

	validation "github.com/go-playground/validator/v10"
)

const invalidJSONPayloadField string = "_payload"
const invalidJSONPayloadTag string = "invalid_json"

type ErrorFormatter struct{}

func NewErrorFormatter() *ErrorFormatter {
	var errorFormatter *ErrorFormatter = &ErrorFormatter{}

	return errorFormatter
}

func (ef *ErrorFormatter) normalizeValidationFieldName(fieldName string) string {
	var fieldNameSegments []string = strings.Split(fieldName, "[")

	return fieldNameSegments[0]
}

func (ef *ErrorFormatter) isJSONPayloadError(validationError error) bool {
	if _, ok := errors.AsType[*json.SyntaxError](validationError); ok {
		return true
	}

	if _, ok := errors.AsType[*json.UnmarshalTypeError](validationError); ok {
		return true
	}

	if errors.Is(validationError, io.EOF) {
		return true
	}

	return false
}

func (ef *ErrorFormatter) FormatValidationErrors(validationError error) map[string]string {
	var formattedValidationErrors map[string]string = map[string]string{}

	if validationError == nil {
		return formattedValidationErrors
	}

	if validationErrors, ok := errors.AsType[validation.ValidationErrors](validationError); ok {
		var validationErrorCount int = len(validationErrors)
		var validationErrorIndex int = 0

		for validationErrorIndex = 0; validationErrorIndex < validationErrorCount; validationErrorIndex++ {
			var fieldValidationError validation.FieldError = validationErrors[validationErrorIndex]
			var fieldName string = ef.normalizeValidationFieldName(fieldValidationError.Field())

			if fieldName == "" {
				fieldName = fieldValidationError.StructField()
			}

			formattedValidationErrors[fieldName] = fieldValidationError.Tag()
		}

		return formattedValidationErrors
	}

	if ef.isJSONPayloadError(validationError) {
		formattedValidationErrors[invalidJSONPayloadField] = invalidJSONPayloadTag

		return formattedValidationErrors
	}

	formattedValidationErrors[invalidJSONPayloadField] = invalidJSONPayloadTag

	return formattedValidationErrors
}
