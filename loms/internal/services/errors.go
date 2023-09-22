package usecase

import "errors"

var (
	// todo add errors and wraps
	ErrReserveStocks = errors.New("reserve stocks failed")
	ErrCreateOrder   = errors.New("create order failed")
)
