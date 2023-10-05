package response

type ErrorDetail struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type ErrorResponse struct {
	Errors []ErrorDetail `json:"errors"`
}
