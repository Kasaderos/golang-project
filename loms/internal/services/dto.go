package usecase

import "route256/loms/internal/models"

type CreateOrderInfo struct {
	Items []models.ItemOrderInfo
}
