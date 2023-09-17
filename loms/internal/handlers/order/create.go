package order

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"route256/loms/internal/models"
	"route256/loms/internal/pkg/handlers"
	"route256/loms/internal/usecase"
)

var ErrNoItems = errors.New("no items")
var ErrIncorrectUser = errors.New("incorrect user")

type CreateReq struct {
	User  int64 `json:"user,omitempty"`
	Items []struct {
		SKU   uint32 `json:"sku,omitempty"`
		Count uint16 `json:"count,omitempty"`
	} `json:"items,omitempty"`
}

func (r CreateReq) Validate() error {
	if len(r.Items) < 1 {
		return ErrNoItems
	}

	if r.User <= 0 {
		return ErrIncorrectUser
	}

	return nil
}

type CreateResp struct {
	OrderID int64
}

type Creator interface {
	Create(ctx context.Context, userID models.UserID, info usecase.CreateOrderInfo) (models.OrderID, error)
}

type CreateHandler struct {
	name    string
	creator Creator
}

func NewCreateHandler(creator Creator) *CreateHandler {
	return &CreateHandler{
		name:    "create order handler",
		creator: creator,
	}
}

func extractOrder(req *CreateReq) (models.UserID, usecase.CreateOrderInfo) {
	// todo
	return models.UserID(0), usecase.CreateOrderInfo{}
}

func (h CreateHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var req CreateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handlers.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		handlers.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	userID, info := extractOrder(&req)

	orderID, err := h.creator.Create(ctx, userID, info)
	if err != nil {
		handlers.GetErrorResponse(w, h.name, err, http.StatusPreconditionFailed)
		return
	}

	resp := CreateResp{
		OrderID: int64(orderID),
	}

	respBytes, err := json.Marshal(resp)
	if err != nil {
		handlers.GetErrorResponse(w, h.name, err, http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(respBytes)

}
