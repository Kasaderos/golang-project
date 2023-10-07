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
	StatusNull
)

var statuses = map[Status]string{
	StatusNew:             "new",
	StatusAwaitingPayment: "awaiting",
	StatusFailed:          "failed",
	StatusPaid:            "paid",
	StatusCancelled:       "cancelled",
	StatusNull:            "null",
}

func (d Status) String() string {
	return statuses[d]
}

func GetStatus(s string) Status {
	switch s {
	case StatusNew.String():
		return StatusNew
	case StatusAwaitingPayment.String():
		return StatusAwaitingPayment
	case StatusFailed.String():
		return StatusFailed
	case StatusPaid.String():
		return StatusPaid
	case StatusCancelled.String():
		return StatusCancelled
	}
	return StatusNull
}
