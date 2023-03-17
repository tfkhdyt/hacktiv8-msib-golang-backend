package errs

import "net/http"

type MessageErr interface {
	Message() string
	StatusCode() int
	Error() string
}

type MessageErrData struct {
	ErrMessage    string `json:"message"`
	ErrStatusCode int    `json:"statusCode"`
	ErrError      string `json:"error"`
}

func (e *MessageErrData) Message() string {
	return e.ErrMessage
}

func (e *MessageErrData) StatusCode() int {
	return e.ErrStatusCode
}

func (e *MessageErrData) Error() string {
	return e.ErrError
}

func NewInternalServerError(message string) MessageErr {
	return &MessageErrData{
		ErrMessage:    message,
		ErrStatusCode: http.StatusInternalServerError,
		ErrError:      "INTERNAL_SERVER_ERROR",
	}
}

func NewUnprocessableEntity(message string) MessageErr {
	return &MessageErrData{
		ErrMessage:    message,
		ErrStatusCode: http.StatusUnprocessableEntity,
		ErrError:      "INVALID_REQUEST_BODY",
	}
}
