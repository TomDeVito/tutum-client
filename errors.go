package tutum

import "errors"

var (
	ErrBadRequest           = errors.New("bad request")
	ErrUnauthorized         = errors.New("unauthorized")
	ErrMethodNotAllowed     = errors.New("method not allowed")
	ErrNotFound             = errors.New("not found")
	ErrUnsupportedMediaType = errors.New("unsupported media type")
	ErrTooManyRequests      = errors.New("too many requests")
	ErrInternalServerError  = errors.New("internal server error")
	ErrServiceUnavailable   = errors.New("service unavailable")
)
