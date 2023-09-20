package controller_http

import (
	"encoding/json"
	"net/http"
	"route256/cart/internal/models"
)

type CheckoutRequest struct {
	User int64 `json:"user,omitempty"`
}

type CheckoutResponse struct {
	OrderID int64 `json:"order_id,omitempty"`
}

func (c *Controller) CheckoutHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	var req CheckoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := req.validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	orderID, err := c.Usecases.Checkout(ctx, models.UserID(req.User))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := CheckoutResponse{
		OrderID: int64(orderID),
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (req *CheckoutRequest) validate() error {
	return nil
}
