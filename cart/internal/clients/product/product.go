package product

import (
	"context"
	"log"
	client_conv "route256/cart/internal/converter/client"
	"route256/cart/internal/models"
	"route256/cart/internal/services/cart"
	products_v1 "route256/cart/pkg/api/products/v1"

	"google.golang.org/grpc/metadata"
)

type contextKeyType string

const tokenContextKey = contextKeyType("token")

const ProductServiceMetadataKey = "Authorization"

type Client struct {
	products_v1.ProductServiceClient
}

var _ cart.ProductProvider = (*Client)(nil)

func NewClient(c products_v1.ProductServiceClient) *Client {
	return &Client{
		ProductServiceClient: c,
	}
}

func WithToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, tokenContextKey, token)
}

func getTokenFromContext(ctx context.Context) string {
	val, ok := ctx.Value(tokenContextKey).(string)
	if !ok || len(val) < 1 {
		log.Println("bad token")
	}
	return val
}

func (c *Client) GetProductInfo(ctx context.Context, sku models.SKU) (name string, price uint32, err error) {
	md := metadata.New(nil)
	md.Set(ProductServiceMetadataKey, getTokenFromContext(ctx))
	ctx = metadata.NewOutgoingContext(ctx, md)

	req := client_conv.ToGetProductRequest(sku)

	resp, err := c.ProductServiceClient.GetProduct(ctx, req)
	if err != nil {
		return "", 0, err
	}

	return resp.Name, uint32(resp.Price), nil
}
