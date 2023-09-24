package repository

import (
	"route256/loms/internal/models"
	"time"
)

type OrderID int64

type UserID int64

type SKU uint32

type Order struct {
	ID        OrderID
	UserID    UserID
	Status    string
	Items     []ItemOrderInfo
	CreatedAt time.Time
}

func (o Order) Scan(*models.Order) {
	// todo
}

type ItemOrderInfo struct {
	SKU   SKU
	Count uint16
}
