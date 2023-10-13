package models

type Cart struct {
	UserID UserID
	Items  []CartItem
}

type CartItem struct {
	SKU   SKU
	Name  string
	Count uint16
	Price uint32
}

type StockItem struct {
	Count uint16
}
