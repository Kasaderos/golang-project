package services

var ErrNotEnoughStocks = NewCartServiceError("not enough stocks")

type CartServiceError struct {
	Message string
}

func NewCartServiceError(msg string) *CartServiceError {
	return &CartServiceError{
		Message: msg,
	}
}

func (c CartServiceError) Error() string {
	return c.Message
}
