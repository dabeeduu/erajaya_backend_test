package customerror

import (
	"backend_golang/utils/errormessage"
	"net/http"
)

const (
	ERRPRODREPOGETALLPROD     = "311GETALLPROD"
	ERRPRODUSECASEGETALLPROD  = "210GETALLPROD"
	ERRPRODHANDLERLISTALLPROD = "110LISTALLPROD"
	ERRPRODREPOINPROD         = "312INPROD"
	ERRUSECASEADDPROD         = "210ADDPROD"
	ERRPRODHANDLERADDPROD     = "110ADDPROD"
	ERRPRODHANDLERADDPRODBIND = "110ADDPRODB"
)

type CustomError struct {
	Code    string
	Message string
	Err     error
}

func (e *CustomError) Codes() string {
	if inner, ok := e.Err.(*CustomError); ok {
		if ToHttpStatus(e.Code) != http.StatusInternalServerError {
			return e.Code
		}
		return inner.Codes()
	}
	return e.Code
}

func (e *CustomError) Error() string {
	if inner, ok := e.Err.(*CustomError); ok {
		return e.Code + ":" + e.Message + "\n" + inner.Error()
	}

	if e.Err != nil {
		return e.Code + ":" + e.Message + " : " + e.Err.Error()
	}

	return e.Code + ":" + e.Message
}

func New(code string, message string, err error) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func NewWithLastCustomError(code string, err error) *CustomError {
	cErr, ok := err.(*CustomError)
	if !ok {
		return &CustomError{
			Code:    code,
			Message: errormessage.ErrorInternalError,
			Err:     err,
		}
	}

	message := cErr.Message
	if message == "" {
		message = errormessage.ErrorInternalError
	}

	return &CustomError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func ToHttpStatus(code string) int {
	switch code {
	default:
		return http.StatusInternalServerError
	}
}

func (e *CustomError) UnWrap() error {
	return e.Err
}
