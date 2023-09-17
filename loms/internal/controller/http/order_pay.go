package controller_http

import (
	"encoding/json"
	"net/http"
	"route256/loms/internal/models"
)

type OrderPayRequest struct {
	OrderID int64 `json:"orderID,omitempty"`
}

func (c *Controller) OrderPayHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	var req OrderPayRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := req.validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.OrderManagementSystem.MarkOrderAsPaid(
		ctx,
		models.OrderID(req.OrderID),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
}

func (req *OrderPayRequest) validate() error {
	return nil
}
