package errno

// Errno -
type Errno struct {
	Code    int
	Message string
	Errors  interface{}
}

// New -
func New(err *Errno, errors interface{}) *Errno {
	e := &Errno{
		Code:    err.Code,
		Message: err.Message,
	}
	if errors != nil {
		e.Errors = errors
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

	return InternalServerError.Code, InternalServerError.Message, err.Error()
}
