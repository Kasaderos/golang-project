package controller_http

import "route256/cart/internal/services"

type Services struct {
	services.CartService
}

type Controller struct {
	ItemAddService
	CheckoutService
	ItemDeleteService
	ListItemService
	ClearService
}

func NewController(us Services) *Controller {
	return &Controller{}
}
