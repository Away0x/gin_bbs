package errno

import (
	"fmt"
)

// Errno -
type Errno struct {
	Code    int
	Message string
	Errors  interface{}
}

// New -
func New(err *Errno, es interface{}) *Errno {
	e := &Errno{
		Code:    err.Code,
		Message: err.Message,
	}
	if es != nil {
		switch typed := es.(type) {
		case *Errno:
			if typed != nil {
				e.Errors = typed.Message
			}
		case error:
			e.Errors = typed.Error()
		default:
			e.Errors = es
		}
	}
	return e
}

// Base -
func Base(err *Errno, msg string) *Errno {
	e := &Errno{
		Code:    err.Code,
		Message: err.Message,
	}

	if msg != "" {
		e.Message = msg
	}

	return e
}

// Error -
func (err Errno) Error() string {
	return err.Message
}

// Decode -
func Decode(err error) (int, string, interface{}) {
	if err == nil {
		return OK.Code, OK.Message, nil
	}

	switch typed := err.(type) {
	case *Errno:
		return typed.Code, typed.Message, typed.Errors
	default:
	}

	fmt.Println(err.Error())

	return InternalServerError.Code, InternalServerError.Message, err.Error()
}
