package domain

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
} //@name Response

type ErrorResponse struct {
	Code         string        `json:"code"`
	ErrorDetails []ErrorDetail `json:"error_details,omitempty" swaggerignore:"true"`
	DebugMessage string        `json:"debug_message,omitempty"`
} //@name ErrorResponse

type ErrorDetail struct {
	Field        string `json:"field"`
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
} //@name ErrorDetail
