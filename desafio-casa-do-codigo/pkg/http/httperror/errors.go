package httperror

import (
	"encoding/json"
	"net/http"
)

type HttpError struct {
	Err        error `json:"-"`
	StatusCode int   `json:"-"`

	StatusText string `json:"status"`
	ErrorText  string `json:"error"`
}

func NewHttpError(status int, err error) HttpError {
	return HttpError{
		Err:        err,
		StatusCode: status,

		StatusText: http.StatusText(status),
		ErrorText:  err.Error(),
	}
}

func ResponseError(w http.ResponseWriter, err HttpError) {
	w.WriteHeader(err.StatusCode)
	_ = json.NewEncoder(w).Encode(err)
}

func BadRequest(w http.ResponseWriter, err error) {
	ResponseError(w, NewHttpError(http.StatusBadRequest, err))
}

func InternalServerError(w http.ResponseWriter, err error) {
	ResponseError(w, NewHttpError(http.StatusInternalServerError, err))
}

func Conflict(w http.ResponseWriter, err error) {
	ResponseError(w, NewHttpError(http.StatusConflict, err))
}
