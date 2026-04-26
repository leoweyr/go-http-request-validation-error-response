package response

type ErrorBody struct {
	Message string            `json:"message"`
	Details map[string]string `json:"details"`
}
