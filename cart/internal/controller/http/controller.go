package controller_http

type Services struct {
	ItemAddService
	CheckoutService
	ItemDeleteService
	ListItemService
	ClearService
}

type Controller struct {
	itemAddService    ItemAddService
	checkoutService   CheckoutService
	itemDeleteService ItemDeleteService
	listItemService   ListItemService
	clearService      ClearService
}

func NewController(srvs Services) *Controller {
	return &Controller{
		itemAddService:    srvs.ItemAddService,
		checkoutService:   srvs.CheckoutService,
		itemDeleteService: srvs.ItemDeleteService,
		listItemService:   srvs.ListItemService,
		clearService:      srvs.ClearService,
	}
}
