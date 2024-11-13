package errList

import "fmt"

const (
	ErrorCategorySystem int8 = iota
	ErrorCategoryBadRequest
	ErrorCategoryUnauthorized

	EventErrorTypeOK int8 = iota + 100
	EventErrorTypeRetry
	EventErrorTypeSkip
	EventErrorTypeUnprocessable
)

// FIXME : error should implement the builtin error interface and should be type nil when it is empty
type CoreError struct {
	Code           int32
	Category       int8
	EventErrorType int8
	ErrorType      string
	Description    string
}

func (e CoreError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Description)
}

// TODO : there might be a better way to handle this
var ErrList = map[eCode]CoreError{
	C1000: {Code: 1000, Category: ErrorCategorySystem, Description: "System error occurred. Please retry."},
	C1001: {Code: 1001, Category: ErrorCategoryBadRequest, EventErrorType: EventErrorTypeSkip, Description: "Invalid request data."},
}

type eCode func()

func C1000() {
}            //System error occurred. Please retry.
func C1001() {} //Invalid request data.

func NewErrorByErrorCode(eCode eCode) CoreError {
	return ErrList[eCode]
}
