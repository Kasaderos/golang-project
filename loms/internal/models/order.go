package models

type Order struct {
	ID     OrderID
	UserID UserID
	Status string
	Items  []ItemOrderInfo
}

type ItemOrderInfo struct {
	SKU   SKU
	Count uint16
}
