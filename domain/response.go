package domain

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

type ErrorResponse struct {
	Code         string        `json:"code"`
	ErrorDetails []ErrorDetail `json:"error_details,omitempty"`
	DebugMessage string        `json:"debug_message,omitempty"`
}

type ErrorDetail struct {
	Field        string `json:"field"`
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}
