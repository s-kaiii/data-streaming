package errorHandling

type ErrorMapping struct {
	Code           int32
	Category       int8
	EventErrorType int8
	ErrorType      string
	Description    string
}

var ErrList = map[int32]ErrorMapping{
	1000: {Code: 1000, Category: ErrorCategorySystem, Description: "System error occurred. Please retry."},
	1001: {Code: 1001, Category: ErrorCategoryBadRequest, EventErrorType: EventErrorTypeSkip, Description: "Invalid request data."},
}
