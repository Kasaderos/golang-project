package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type StocksService interface {
	GetStocks(sku uint32) uint64
}

type StocksHandler struct {
	stocksService StocksService
}

func NewStocksHandler(stocksService StocksService) *StocksHandler {
	return &StocksHandler{
		stocksService: stocksService,
	}
}

type StocksRequest struct {
	SKU uint32 `json:"sku,omitempty"`
}

var ErrIncorrectSKU = errors.New("incorrect SKU")

func (r StocksRequest) Validate() error {
	if r.SKU == 0 {
		return ErrIncorrectSKU
	}
	return nil
}

type StocksResponse struct {
	Count uint64 `json:"count,omitempty"`
}

func (s StocksHandler) Handle(w http.ResponseWriter, r *http.Request) {
	req := &StocksRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Println("failed to decode request body")
		GetErrorResponse(w, "stocks", err, http.StatusBadRequest)
		return
	}
	if err := req.Validate(); err != nil {
		log.Println("stocks: failed to validate request:", err)
		GetErrorResponse(w, "stocks", err, http.StatusBadRequest)
		return
	}

	count := s.stocksService.GetStocks(req.SKU)

	stocksResponse := &StocksResponse{
		Count: count,
	}
	raw, err := json.Marshal(stocksResponse)
	if err != nil {
		GetErrorResponse(w, "stocks", err, http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	GetSuccessResponseWithBody(w, raw)
}
