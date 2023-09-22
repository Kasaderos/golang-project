package controller_http

import (
	"context"
	"encoding/json"
	"net/http"
	"route256/loms/internal/models"
	dto "route256/loms/internal/services"
)

type OrderCreateService interface {
	CreateOrder(ctx context.Context, userID models.UserID, info dto.CreateOrderInfo) (models.OrderID, error)
}

type OrderCreateRequest struct {
	UserID int64 `json:"user"`
	Items  []struct {
		SKU   int64  `json:"sku"`
		Count uint16 `json:"count"`
	} `json:"items,omitempty"`
}

type OrderCreateResponse struct {
	OrderID int64 `json:"orderID"`
}

func (c *Controller) OrderCreateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	var req OrderCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := req.validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	orderID, err := c.orderCreateService.CreateOrder(
		ctx,
		models.UserID(req.UserID),
		req.CreateOrderInfo(),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// according to the task OK
	// but I think we should return 201
	w.WriteHeader(http.StatusOK)

	resp := OrderCreateResponse{
		OrderID: int64(orderID),
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (req *OrderCreateRequest) validate() error {
	return nil
}

func (req *OrderCreateRequest) CreateOrderInfo() dto.CreateOrderInfo {
	info := dto.CreateOrderInfo{
		Items: make([]models.ItemOrderInfo, 0, len(req.Items)),
	}

	for _, item := range req.Items {
		info.Items = append(info.Items, models.ItemOrderInfo{
			SKU:   models.SKU(item.SKU),
			Count: item.Count,
		})
	}

	return info
}
