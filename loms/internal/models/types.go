package models

type OrderID int64

type UserID int64

type SKU uint32

type Status uint32

const (
	StatusNew = iota + 1
	StatusAwaitingPayment
	StatusFailed
	StatusPaid
	StatusCancelled
)
