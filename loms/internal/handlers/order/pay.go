package order

import (
	"context"
	"encoding/json"
	"net/http"
	"route256/loms/internal/pkg/handlers"
)

type PayReq struct {
	OrderID int64 `json:"order_id,omitempty"`
}

func (r PayReq) Validate() error {
	// if r.OrderID <= 0 {
	// 	return ErrIncorrectOrderID
	// }
	// // todo
	// // check if it exists

	return nil
}

type PayMarker interface {
	MarkAsPaid(ctx context.Context, orderID int64) error
}

type PayHandler struct {
	name    string
	payment PayMarker
}

func NewPayHandler(payment PayMarker) *PayHandler {
	return &PayHandler{
		name:    "order pay handler",
		payment: payment,
	}
}

func (h PayHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var req PayReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handlers.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		handlers.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	err := h.payment.MarkAsPaid(ctx, req.OrderID)
	if err != nil {
		// todo 500
		handlers.GetErrorResponse(w, h.name, err, http.StatusNotFound)
		return
	}

	_, _ = w.Write([]byte("{}"))
}
