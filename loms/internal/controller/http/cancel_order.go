package controller_http

import (
	"context"
	"encoding/json"
	"net/http"
	"route256/loms/internal/models"
)

type Stocks interface {
	GetStockInfo(
		ctx context.Context,
		SKU models.SKU,
	) (count uint64, err error)
}

type OrderCanceller interface {
	CancelOrder(
		ctx context.Context,
		orderID models.OrderID,
	) error
}

type CancelOrderRequest struct {
	OrderID int64 `json:"orderID,omitempty"`
}

func (c *Controller) CancelOrderHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	var req CancelOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := req.validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.OrderCanceller.CancelOrder(
		ctx,
		models.OrderID(req.OrderID),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
}

func (req *CancelOrderRequest) validate() error {
	return nil
}
