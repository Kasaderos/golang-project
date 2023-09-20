package controller_http

import (
	"encoding/json"
	"net/http"
	"route256/cart/internal/models"
)

type ItemAddRequest struct {
	User  int64  `json:"user,omitempty"`
	SKU   uint32 `json:"sku,omitempty"`
	Count uint16 `json:"count,omitempty"`
}

func (c *Controller) ItemAddHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusNotFound)
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

	err := c.Usecases.AddItem(
		ctx,
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
	return nil
}
