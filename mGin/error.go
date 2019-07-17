package mGin

import "net/http"

// Error 带有错误码 与 错误信息的错误类
type Error interface {
	error
	Code() int
}

// ErrNew returns an error that formats as the given text.
func ErrNew(code int, text string) Error {
	return &errorString{code, text}
}

// ErrNew returns an error that formats as the given text.
func ErrNewErr(err error) Error {
	return &errorString{http.StatusInternalServerError, err.Error(),}
}

// errorString is a trivial implementation of error.
type errorString struct {
	code int
	s    string
}

func (e *errorString) Error() string {
	return e.s
}

func (e *errorString) Code() int {
	return e.code
}
