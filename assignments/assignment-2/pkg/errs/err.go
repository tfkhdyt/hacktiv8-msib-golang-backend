package errs

import "net/http"

type MessageErr interface {
	Message() string
	StatusCode() int
	Error() string
}

type messageErrData struct {
	ErrMessage    string `json:"message"`
	ErrStatusCode int    `json:"statusCode"`
	ErrError      string `json:"error"`
}

func (e *messageErrData) Message() string {
	return e.ErrMessage
}

func (e *messageErrData) StatusCode() int {
	return e.ErrStatusCode
}

func (e *messageErrData) Error() string {
	return e.ErrError
}

func NewInternalServerError(message string) MessageErr {
	return &messageErrData{
		ErrMessage:    message,
		ErrStatusCode: http.StatusInternalServerError,
		ErrError:      "INTERNAL_SERVER_ERROR",
	}
}

func NewUnprocessableEntity(message string) MessageErr {
	return &messageErrData{
		ErrMessage:    message,
		ErrStatusCode: http.StatusUnprocessableEntity,
		ErrError:      "INVALID_REQUEST_BODY",
	}
}

func NewBadRequest(message string) MessageErr {
	return &messageErrData{
		ErrMessage:    message,
		ErrStatusCode: http.StatusBadRequest,
		ErrError:      "BAD_REQUEST",
	}
}
