package loms

import (
	"context"
	client_conv "route256/cart/internal/converter/client"
	"route256/cart/internal/models"
	loms_v1 "route256/loms/pkg/api/loms/v1"
)

type Client struct {
	loms_v1.LOMSClient
}

func NewClient(c loms_v1.LOMSClient) *Client {
	return &Client{LOMSClient: c}
}

func (s *Client) CreateOrder(
	ctx context.Context,
	userID models.UserID,
	items []models.CartItem,
) (models.OrderID, error) {
	resp, err := s.LOMSClient.OrderCreate(ctx, client_conv.ToOrderCreateRequest(userID, items))
	if err != nil {
		return models.OrderID(0), err
	}

	return models.OrderID(resp.OrderId), nil
}

func (c *Client) GetStock(ctx context.Context, sku models.SKU) (count uint64, err error) {
	resp, err := c.LOMSClient.GetStockInfo(ctx, client_conv.ToGetStockRequest(sku))
	if err != nil {
		return 0, err
	}

	return resp.Count, nil
}
