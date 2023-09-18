package models

type Cart struct {
	UserID UserID
	Items  []CartItem
}

type CartItem struct {
	SKU   SKU
	Count uint16
}

type StockItem struct {
	Count uint16
}
