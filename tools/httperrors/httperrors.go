package httperrors

import (
	"net/http"

	"github.com/go-chi/render"
)

type HttpResponseError struct {
	Err error `json:"-"` // low-level runtime error

	Status  int    `json:"status"`  // http status code
	Code    string `json:"code"`    // application-specific error code
	Message string `json:"message"` // user-level message
}

func (e *HttpResponseError) Error() string {
	return e.Message
}

func (e *HttpResponseError) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.Status)

	return nil
}

func NewBadRequestError(err error) render.Renderer {
	return &HttpResponseError{
		Err:     err,
		Status:  http.StatusBadRequest,
		Code:    "bad_request",
		Message: err.Error(),
	}
}

func NewNotFoundError(err error) render.Renderer {
	return &HttpResponseError{
		Err:     err,
		Status:  http.StatusNotFound,
		Code:    "not_found",
		Message: err.Error(),
	}
}

func NewInternalServerError(err error) render.Renderer {
	return &HttpResponseError{
		Err:     err,
		Status:  http.StatusInternalServerError,
		Code:    "internal_server_error",
		Message: err.Error(),
	}
}

func NewUnauthorizedError(err error) render.Renderer {
	return &HttpResponseError{
		Err:     err,
		Status:  http.StatusUnauthorized,
		Code:    "unauthorized",
		Message: err.Error(),
	}
}

func NewForbiddenError(err error) render.Renderer {
	return &HttpResponseError{
		Err:     err,
		Status:  http.StatusForbidden,
		Code:    "forbidden",
		Message: err.Error(),
	}
}

func NewConflictError(err error) render.Renderer {
	return &HttpResponseError{
		Err:     err,
		Status:  http.StatusConflict,
		Code:    "conflict",
		Message: err.Error(),
	}
}
