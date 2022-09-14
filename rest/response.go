package response

type Response struct {
	Success     bool        `json:"success"`
	Message     string      `json:"message"`
	ProcessTime string      `json:"process_time"`
	ErrorCode   string      `json:"error_code,omitempty"`
	Data        interface{} `json:"data"`
}
