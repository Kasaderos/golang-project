package carts

import (
	servicepb "route256/cart/pkg/api/carts/v1"
)

type Deps struct {
	ItemAddService
	CheckoutService
	ItemDeleteService
	ListItemService
	ClearService
}

type Service struct {
	servicepb.UnimplementedCartsServer
	itemAddService    ItemAddService
	checkoutService   CheckoutService
	itemDeleteService ItemDeleteService
	listItemService   ListItemService
	clearService      ClearService
}

func NewServer(d *Deps) *Service {
	return &Service{
		itemAddService:    d.ItemAddService,
		checkoutService:   d.CheckoutService,
		itemDeleteService: d.ItemDeleteService,
		listItemService:   d.ListItemService,
		clearService:      d.ClearService,
	}
}
