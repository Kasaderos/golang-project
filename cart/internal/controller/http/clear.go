package controller_http

import (
	"context"
	"encoding/json"
	"net/http"
	"route256/cart/internal/models"
)

type ClearService interface {
	Clear(ctx context.Context, userID models.UserID) error
}

type ClearRequest struct {
	User int64 `json:"user,omitempty"`
}

func (c *Controller) ClearHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	var req ClearRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := req.validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.ClearService.Clear(ctx, models.UserID(req.User))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (req *ClearRequest) validate() error {
	return nil
}
