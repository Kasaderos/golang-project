package repository

import (
	"context"
	"route256/loms/internal/models"
	oms "route256/loms/internal/usecase/OMS"
)

type omsRepository struct {
}

var _ oms.OMSRepository = (*omsRepository)(nil)

func NewOMSRepostiory() *omsRepository {
	return &omsRepository{}
}

func (r *omsRepository) CreateOrder(ctx context.Context, order models.Order) (models.Order, error) {

	return models.Order{}, nil
}
