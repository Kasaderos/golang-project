package product

import (
	"route256/cart/internal/models"
	products_v1 "route256/cart/pkg/api/products/v1"
)

func ToGetProductRequest(sku models.SKU, token string) *products_v1.GetProductRequest {
	return &products_v1.GetProductRequest{
		Sku:   uint32(sku),
		Token: token,
	}
}
