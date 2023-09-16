package models

type Order struct {
	User  int64
	Items []OrderItem
}

type OrderItem struct {
	SKU   uint32
	Count uint16
}
