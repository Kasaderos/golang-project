package loms

import (
	"context"
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
	// todo
	// add converter pkg
	reqItems := make([]*loms_v1.OrderInfoItem, 0, len(items))
	for _, item := range items {
		reqItems = append(reqItems, &loms_v1.OrderInfoItem{
			Sku:   int64(item.Count),
			Count: uint32(item.Count),
		})
	}
	in := &loms_v1.OrderCreateRequest{
		User:  int64(userID),
		Items: reqItems,
	}
	resp, err := s.LOMSClient.OrderCreate(ctx, in)
	if err != nil {
		return models.OrderID(0), err
	}

	return models.OrderID(resp.OrderId), nil
}

func (c *Client) GetStock(ctx context.Context, sku models.SKU) (count uint64, err error) {
	// todo
	// add converter pkg
	req := &loms_v1.GetStockInfoRequest{
		Sku: uint32(sku),
	}
	resp, err := c.LOMSClient.GetStockInfo(ctx, req)
	if err != nil {
		return 0, err
	}

	return resp.Count, nil
}
