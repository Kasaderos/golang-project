package controller_http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"route256/cart/internal/clients/product"
	"route256/cart/internal/models"
)

var ErrCountIsZero = errors.New("count is zero")

type ItemAddService interface {
	AddItem(ctx context.Context, userID models.UserID, sku models.SKU, count uint16) error
}

type ItemAddRequest struct {
	User  int64  `json:"user,omitempty"`
	SKU   uint32 `json:"sku,omitempty"`
	Count uint16 `json:"count,omitempty"`
}

func (c *Controller) ItemAddHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	productServiceToken := r.Header.Get("X-Product-Service-Token")
	if len(productServiceToken) < 1 {
		http.Error(w, "empty product service token", http.StatusBadRequest)
		return
	}

	var req ItemAddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := req.validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.itemAddService.AddItem(
		product.WithToken(ctx, productServiceToken),
		models.UserID(req.User),
		models.SKU(req.SKU),
		req.Count,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}
}

func (req *ItemAddRequest) validate() error {
	if req.Count <= 0 {
		return ErrCountIsZero
	}
	return nil
}
