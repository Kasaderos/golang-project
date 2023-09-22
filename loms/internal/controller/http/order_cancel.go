package controller_http

import (
	"context"
	"encoding/json"
	"net/http"
	"route256/loms/internal/models"
)

type StocksService interface {
	GetStockInfo(
		ctx context.Context,
		SKU models.SKU,
	) (count uint64, err error)
}

type OrderCancelService interface {
	CancelOrder(
		ctx context.Context,
		orderID models.OrderID,
	) error
}

type OrderCancelRequest struct {
	OrderID int64 `json:"orderID,omitempty"`
}

func (c *Controller) OrderCancelHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req OrderCancelRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := req.validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.orderCancelService.CancelOrder(
		ctx,
		models.OrderID(req.OrderID),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
}

func (req *OrderCancelRequest) validate() error {
	return nil
}
