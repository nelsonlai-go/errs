package errs

import "fmt"

type Error interface {
	error
	Code() string
	FullMessage() string
}

type err struct {
	code    string
	message string
}

func (e *err) Error() string {
	return e.message
}

func (e *err) Code() string {
	return e.code
}

func (e *err) FullMessage() string {
	return fmt.Sprintf("[%s] %s", e.code, e.message)
}

func New(code, message string) Error {
	return &err{code: code, message: message}
}
