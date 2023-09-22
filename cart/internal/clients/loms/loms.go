package loms

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"route256/cart/internal/models"
	"route256/cart/internal/services/cart"
)

var (
	ErrOrderNotCreated = errors.New("order not created")
)

const (
	CreateOrderAPIPath = "/order/create"
	GetStockAPIPath    = "/stock/info"
)

type lomsService struct {
	name       string
	baseURL    string
	httpClient *http.Client
}

var _ cart.OrderCreator = (*lomsService)(nil)

func NewLOMSService(baseURL string) *lomsService {
	return &lomsService{
		name:       "loms",
		httpClient: &http.Client{},
		baseURL:    baseURL,
	}
}

func (srv *lomsService) CreateOrder(
	ctx context.Context,
	userID models.UserID,
	items []models.CartItem,
) (models.OrderID, error) {
	body := CreateOrderRequest{
		UserID: int64(userID),
		Items:  make([]CreateOrderItem, 0, len(items)),
	}
	for _, item := range items {
		body.Items = append(body.Items, CreateOrderItem{
			SKU:   int64(item.SKU),
			Count: item.Count,
		})
	}

	reqBody, err := json.Marshal(body)
	if err != nil {
		return models.OrderID(0), err
	}

	reqURL, err := url.JoinPath(srv.baseURL, CreateOrderAPIPath)
	if err != nil {
		return models.OrderID(0), err
	}

	req, err := http.NewRequest(http.MethodPost, reqURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return models.OrderID(0), err
	}

	resp, err := srv.httpClient.Do(req)
	if err != nil {
		return models.OrderID(0), err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != 200 {
		var response CreateOrderErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			return models.OrderID(0), err
		}
		return models.OrderID(0), fmt.Errorf("%s: responded with: %s", srv.name, response.Message)
	}

	var respSt CreateOrderResponse
	if err := json.NewDecoder(resp.Body).Decode(&resp); err != nil {
		return models.OrderID(0), fmt.Errorf("%s: decode: %w", srv.name, err)
	}

	return models.OrderID(respSt.OrderID), nil
}

func (srv *lomsService) GetStock(ctx context.Context, sku models.SKU) (count uint64, err error) {
	body := GetStockInfoRequest{
		SKU: uint32(sku),
	}

	reqBody, err := json.Marshal(body)
	if err != nil {
		return 0, err
	}

	reqURL, err := url.JoinPath(srv.baseURL, GetStockAPIPath)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest(http.MethodPost, reqURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return 0, err
	}

	resp, err := srv.httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != 200 {
		var response CreateOrderErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			return 0, err
		}
		return 0, fmt.Errorf("%s: responded with: %s", srv.name, response.Message)
	}

	var stockInfo GetStockInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&stockInfo); err != nil {
		return 0, err
	}

	return stockInfo.Count, nil
}
