package dtos

type Response struct {
	StatusCode int    `json:"status_code"`
	Error      string `json:"error,omitempty"`
	Value      any    `json:"value,omitempty"`
}
