package controller_http

import (
	"encoding/json"
	"net/http"
	"route256/loms/internal/models"
)

type GetStockInfoRequest struct {
	SKU uint32 `json:"sku,omitempty"`
}

type GetStockInfoResponse struct {
	Count uint64 `json:"count,omitempty"`
}

func (c *Controller) StockInfoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	var req GetStockInfoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := req.validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	count, err := c.WMSUsecase.GetStockInfo(
		ctx,
		models.SKU(req.SKU),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resp := GetStockInfoResponse{
		Count: count,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (req *GetStockInfoRequest) validate() error {
	return nil
}
