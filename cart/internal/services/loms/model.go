package loms

type CreateOrderRequest struct {
	UserID int64             `json:"user"`
	Items  []CreateOrderItem `json:"items,omitempty"`
}

type CreateOrderItem struct {
	SKU   int64  `json:"sku"`
	Count uint16 `json:"count"`
}

type CreateOrderResponse struct {
	OrderID int64 `json:"orderID"`
}

type GetStockInfoRequest struct {
	SKU uint32 `json:"sku,omitempty"`
}

type GetStockInfoResponse struct {
	Count uint64 `json:"count,omitempty"`
}

type CreateOrderErrorResponse struct {
	Message string `json:"message"`
}

type GetStockInfoErrorResponse struct {
	Message string `json:"message"`
}
