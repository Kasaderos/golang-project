package controller_http

import "route256/loms/internal/usecase"

type Usecases struct {
	OMSUsecase usecase.OrderManagementSystem
	WMSUsecase usecase.WarehouseManagementSystem
}

type Controller struct {
	Usecases
}

func NewController(us Usecases) *Controller {
	return &Controller{
		Usecases: us,
	}
}
