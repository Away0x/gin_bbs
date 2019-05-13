package errno

type Errno struct {
	Code    int
	Message string
	Errors  interface{}
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
