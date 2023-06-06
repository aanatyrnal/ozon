package httpErrors

import (
	"errors"
	"fmt"
	"net/http"
)

const (
	ErrLinkAlreadyExists = "User with given email already exists"
	ErrNotFound          = "Not Found"
)

var (
	NotFound        = errors.New("Not Found")
	ExistsLinkError = errors.New("Link already exists")
)

type RestErr interface {
	Status() int
	Error() string
	Causes() interface{}
}

type RestError struct {
	ErrStatus int         `json:"status,omitempty"`
	ErrError  string      `json:"error,omitempty"`
	ErrCauses interface{} `json:"-"`
}

// Error  Error() interface method
func (e RestError) Error() string {
	return fmt.Sprintf("status: %d - errors: %s - causes: %v", e.ErrStatus, e.ErrError, e.ErrCauses)
}

func (e RestError) Status() int {
	return e.ErrStatus
}

func (e RestError) Causes() interface{} {
	return e.ErrCauses
}

func NewNotFoundError(causes interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusNotFound,
		ErrError:  ErrNotFound,
		ErrCauses: causes,
	}
}

func NewExistsLinkError(causes interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusBadRequest,
		ErrError:  ErrLinkAlreadyExists,
		ErrCauses: causes,
	}
}
