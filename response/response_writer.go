package response

import (
	"encoding/json"
	"net/http"
)

type ResponseWriter struct{}

func NewResponseWriter() *ResponseWriter {
	var responseWriter *ResponseWriter = &ResponseWriter{}

	return responseWriter
}

func (responseWriter *ResponseWriter) WriteJSON(httpResponseWriter http.ResponseWriter, statusCode int, responseBody any) error {
	httpResponseWriter.Header().Set("Content-Type", "application/json")
	httpResponseWriter.WriteHeader(statusCode)

	var jsonEncoder *json.Encoder = json.NewEncoder(httpResponseWriter)
	var encodeError error = jsonEncoder.Encode(responseBody)

	return encodeError
}
