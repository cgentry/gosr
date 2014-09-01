package gosr

import (
	"errors"
	"net/http"
)

const (
	StatusTokenExpired 	= 498
	StatusTokenRequired = 499
)

// These messages are to override the standard messages
var ResponseMessageExtended = map[int]string {
	http.StatusContinue		:	"OK",
	http.StatusCreated		:	"OK; data saved",
	http.StatusNoContent	:	"OK; but no content returned",

	http.StatusPaymentRequired :"Account not active",
	http.StatusForbidden 	:	"You do not have access permission" ,
	http.StatusNotFound		:	"Action is not implemented",
	http.StatusMethodNotAllowed :	"Method not allowed for operation",

	498 :	"Token has expired",
	499 :	"Token required",
}

type Error struct {
	Code	int
	Text	string
}

func GetErrorText( code int ) string {
	if msg,ok := ResponseMessageExtended[code]; ok {
		return msg
	}
	return http.StatusText(code)
}

func NewError( code int ) * Error {
	return &Error{ Code : code , Text : GetErrorText( code ) }
}

func NewErrorWithText( code  int , msg string ) * Error {
	return &Error{ Code : code , Text : msg }
}

func ( e * Error ) ToError() error {
	if e != nil {
		return errors.New( e.Text )
	}
	return nil
}

func ( e * Error ) Error() string {
	return e.Text
}

func ( e * Error ) Status() int {
	return e.Code
}

