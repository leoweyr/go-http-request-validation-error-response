package request

import (
	"encoding/json"
	"net/http"
)

type RequestReader struct{}

func NewRequestReader() *RequestReader {
	var requestReader *RequestReader = &RequestReader{}

	return requestReader
}

func (requestReader *RequestReader) ReadJSON(httpRequest *http.Request, payload any) error {
	var jsonDecoder *json.Decoder = json.NewDecoder(httpRequest.Body)

	jsonDecoder.DisallowUnknownFields()

	var decodeError error = jsonDecoder.Decode(payload)

	return decodeError
}
