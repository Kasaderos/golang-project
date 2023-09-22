package controller_http

import (
	"context"
	"encoding/json"
	"net/http"
	"route256/cart/internal/models"
)

type ItemDeleteService interface {
	DeleteItem(ctx context.Context, userID models.UserID, sku models.SKU) error
}

type ItemDeleteRequest struct {
	User int64  `json:"user,omitempty"`
	SKU  uint32 `json:"sku,omitempty"`
}

func (c *Controller) ItemDeleteHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	var req ItemDeleteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := req.validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.ItemDeleteService.DeleteItem(
		ctx,
		models.UserID(req.User),
		models.SKU(req.SKU),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (req *ItemDeleteRequest) validate() error {
	return nil
}
