package wms

import (
	"context"
	"route256/loms/internal/models"
	"route256/loms/internal/usecase"
)

type (
	// warehouse management system repository
	WMSRepository interface {
		GetStockBySKU(ctx context.Context, SKU models.SKU) (count uint64, err error)
	}
)

type Deps struct {
	WMSRepository
}

type wmsUsecase struct {
	Deps
}

var _ usecase.WarehouseManagementSystem = (*wmsUsecase)(nil)

func NewWMSUsecase(d Deps) *wmsUsecase {
	return &wmsUsecase{
		Deps: d,
	}
}

func (usc *wmsUsecase) GetStockInfo(
	ctx context.Context,
	SKU models.SKU,
) (count uint64, err error) {
	return usc.WMSRepository.GetStockBySKU(ctx, SKU)
}
