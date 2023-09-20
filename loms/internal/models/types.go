package models

type OrderID int64

type UserID int64

type SKU uint32

type Status uint32

const (
	StatusNew Status = iota + 1
	StatusAwaitingPayment
	StatusFailed
	StatusPaid
	StatusCancelled
)

var statuses = map[Status]string{
	StatusNew:             "new",
	StatusAwaitingPayment: "awaiting",
	StatusFailed:          "failed",
	StatusPaid:            "paid",
	StatusCancelled:       "cancelled",
}

func (d Status) String() string {
	return statuses[d]
}
