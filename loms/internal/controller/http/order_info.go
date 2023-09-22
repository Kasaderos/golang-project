package controller_http

import (
	"context"
	"encoding/json"
	"net/http"
	"route256/loms/internal/models"
)

type OrderInformer interface {
	GetInfo(ctx context.Context, orderID models.OrderID) (*models.Order, error)
}

type OrderInfoRequest struct {
	OrderID int64 `json:"orderID,omitempty"`
}

type OrderInfoResponse struct {
	Status string          `json:"status,omitempty"`
	User   int64           `json:"user,omitempty"`
	Items  []OrderInfoItem `json:"items,omitempty"`
}

type OrderInfoItem struct {
	SKU   uint32 `json:"sku,omitempty"`
	Count uint16 `json:"count,omitempty"`
}

func (c *Controller) OrderInfoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	var req OrderInfoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := req.validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order, err := c.OrderInformer.GetInfo(
		ctx,
		models.OrderID(req.OrderID),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resp := toOrderInfoResponse(order)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (req *OrderInfoRequest) validate() error {
	return nil
}

func toOrderInfoResponse(order *models.Order) *OrderInfoResponse {
	if order == nil {
		return nil
	}
	resp := new(OrderInfoResponse)

	resp.Status = order.Status.String()
	resp.User = int64(order.UserID)
	resp.Items = make([]OrderInfoItem, 0, len(order.Items))

	for _, item := range order.Items {
		resp.Items = append(resp.Items, OrderInfoItem{
			SKU:   uint32(item.SKU),
			Count: item.Count,
		})
	}

	return resp
}
