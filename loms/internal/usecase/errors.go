package usecase

import "errors"

var (
	ErrReserveStocks = errors.New("reserve stocks failed")
	ErrCreateOrder   = errors.New("create order failed")
)
