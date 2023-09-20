package controller_http

import "route256/cart/internal/usecase"

type Usecases struct {
	usecase.CartService
}

type Controller struct {
	Usecases
}

func NewController(us Usecases) *Controller {
	return &Controller{
		Usecases: us,
	}
}
