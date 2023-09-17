package controller_http

import (
	"encoding/json"
	"net/http"
	"route256/loms/internal/models"
	"route256/loms/internal/usecase"
)

type CreateOrderRequest struct {
	UserID int64 `json:"user"`
	Items  []struct {
		SKU   int64  `json:"sku"`
		Count uint16 `json:"count"`
	} `json:"items,omitempty"`
}

type CreateOrderResponse struct {
	OrderID int64 `json:"orderID"`
}

func (c *Controller) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	// 0. Decode request
	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 1. Validation
	if err := validateCreateOrderRequest(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 2. Transform delivery layer models to Domain/Usecase models
	orderInfo := extractCreateOrderInfoFromCreateOrderRequest(&req)

	// 3. Call usecases
	orderID, err := c.OrderManagementSystem.CreateOrder(ctx, models.UserID(req.UserID), orderInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 4. Prepare answer
	resp := CreateOrderResponse{
		OrderID: int64(orderID),
	}

	// 5. Decode answer & send response
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// according to the task OK but I think we should return 201
	w.WriteHeader(http.StatusOK)
}

func validateCreateOrderRequest(req *CreateOrderRequest) error {
	/* your validation logic here */
	return nil
}

func extractCreateOrderInfoFromCreateOrderRequest(req *CreateOrderRequest) usecase.CreateOrderInfo {
	/* your mapping logic here */

	info := usecase.CreateOrderInfo{
		/* ... */
	}

	return info
}
