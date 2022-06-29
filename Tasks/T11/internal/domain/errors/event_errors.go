package errors

import "errors"

var ErrBadRequest = errors.New("bad request")

var ErrWrongQueryParams = errors.New("wrong query parameters")
