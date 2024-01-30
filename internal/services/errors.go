package services

type ErrorMessages struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func (e *ErrorMessages) Error() string {
	return e.Message
}
