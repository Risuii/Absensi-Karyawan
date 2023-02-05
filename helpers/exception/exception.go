package exception

import "fmt"

var (
	ErrConflicted          = fmt.Errorf("conflicted")
	ErrInternalServer      = fmt.Errorf("internal server error")
	ErrNotFound            = fmt.Errorf("not found error")
	ErrBadRequest          = fmt.Errorf("bad request")
	ErrUnauthorized        = fmt.Errorf("unauthorized")
	ErrNotPremium          = fmt.Errorf("not premium user")
	ErrUnprocessableEntity = fmt.Errorf("error UnprocessableEntity")
)
