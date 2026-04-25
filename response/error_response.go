package response

type ErrorResponse struct {
	Error ErrorBody `json:"error"`
}
