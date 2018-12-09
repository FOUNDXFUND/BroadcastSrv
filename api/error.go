package api

import "net/http"

type Error struct {
	HTTPStatus int
	Code int
	Msg string
}
func newError(status,code int, msg string) *Error{
	return &Error{status, code, msg}
}

var (
	ErrWrongParam = newError(http.StatusBadRequest, -998, "Parameter Error")
)