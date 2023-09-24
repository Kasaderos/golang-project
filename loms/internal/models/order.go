package models

type Order struct {
	ID     OrderID
	UserID UserID
	Status Status
	Items  []ItemOrderInfo
}

type ItemOrderInfo struct {
	SKU   SKU
	Count uint16
}
