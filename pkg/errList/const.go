package errorHandling

const (
	ErrorCategorySystem int8 = iota
	ErrorCategoryBadRequest
	ErrorCategoryUnauthorized

	EventErrorTypeOK int8 = iota + 100
	EventErrorTypeRetry
	EventErrorTypeSkip
	EventErrorTypeUnprocessable
)
