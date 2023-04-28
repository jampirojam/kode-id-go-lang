package response

type ErrorResponse struct {
	Error interface{} `json:"error"`
}
