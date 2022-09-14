package response

type ErrorFieldResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
