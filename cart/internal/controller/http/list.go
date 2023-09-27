package controller_http

import (
	"context"
	"encoding/json"
	"net/http"
	"route256/cart/internal/clients/http/product"
	"route256/cart/internal/models"
)

type ListItemService interface {
	ListItem(ctx context.Context, userID models.UserID) (totalPrice uint32, items []models.CartItem, err error)
}

type ListRequest struct {
	User int64 `json:"user,omitempty"`
}

type ListResponse struct {
	Items      []ListItem `json:"items,omitempty"`
	TotalPrice uint32     `json:"total_price,omitempty"`
}

type ListItem struct {
	SKU   uint32 `json:"sku,omitempty"`
	Count uint16 `json:"count,omitempty"`
	Name  string `json:"name,omitempty"`
	Price uint32 `json:"price,omitempty"`
}

func (c *Controller) ListHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	productServiceToken := r.Header.Get("X-Product-Service-Token")
	if len(productServiceToken) < 1 {
		http.Error(w, "empty product service token", http.StatusBadRequest)
		return
	}

	var req ListRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := req.validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	totalPrice, items, err := c.listItemService.ListItem(
		product.WithToken(ctx, productServiceToken),
		models.UserID(req.User),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}

	resp := toListResponse(items, totalPrice)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (req *ListRequest) validate() error {
	return nil
}

func toListResponse(items []models.CartItem, totalPrice uint32) ListResponse {
	resp := ListResponse{
		Items:      make([]ListItem, 0, len(items)),
		TotalPrice: totalPrice,
	}

	for _, item := range items {
		resp.Items = append(resp.Items, ListItem{
			SKU:   uint32(item.SKU),
			Count: item.Count,
			Name:  item.Name,
			Price: item.Price,
		})
	}

	return resp
}
