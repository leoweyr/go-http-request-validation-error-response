package response

var validationFailedMessage string = "VALIDATION_FAILED"

type ErrorResponseBuilder struct{}

func NewErrorResponseBuilder() *ErrorResponseBuilder {
	var errorResponseBuilder *ErrorResponseBuilder = &ErrorResponseBuilder{}

	return errorResponseBuilder
}

func (erb *ErrorResponseBuilder) buildValidationErrorBody(validationDetails map[string]string) ErrorBody {
	var errorBody ErrorBody = ErrorBody{
		Message: validationFailedMessage,
		Details: validationDetails,
	}

	return errorBody
}

func (erb *ErrorResponseBuilder) BuildValidationFailedErrorResponse(validationDetails map[string]string) ErrorResponse {
	var validationErrorResponse ErrorResponse = ErrorResponse{
		Error: erb.buildValidationErrorBody(validationDetails),
	}

	return validationErrorResponse
}
