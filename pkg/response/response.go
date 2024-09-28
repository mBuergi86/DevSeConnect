package response

import "fmt"

type Error struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Err        error  `json:"error"`
}

func NewError(message string, statusCode int, err error) *Error {
	return &Error{
		Message:    message,
		StatusCode: statusCode,
		Err:        err,
	}
}

func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s %v", e.Message, e.Err)
	}
	return e.Message
}
