package orders

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"route256/loms/internal/pkg/handlers"
	"route256/loms/internal/pkg/models"
)

var ErrNoItems = errors.New("no items")
var ErrIncorrectUser = errors.New("incorrect user")

type CreateReq struct {
	User  int64 `json:"user,omitempty"`
	Items []struct {
		SKU   uint32 `json:"sku,omitempty"`
		Count uint16 `json:"count,omitempty"`
	}
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
	Create(ctx context.Context, order *models.Order) (orderID string, err error)
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

func extractOrder(req *CreateReq) *models.Order {
	order := &models.Order{
		User:  req.User,
		Items: make([]models.OrderItem, 0, len(req.Items)),
	}
	for _, item := range req.Items {
		order.Items = append(order.Items, models.OrderItem{
			SKU:   item.SKU,
			Count: item.Count,
		})
	}
	return order
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
	order := extractOrder(&req)

	orderID, err := h.creator.Create(ctx, order)
	if err != nil {
		// todo
		// check server error (500) using "errors" pkg
		handlers.GetErrorResponse(w, h.name, err, http.StatusPreconditionFailed)
		return
	}

	// w.Write()
	_ = orderID

}
