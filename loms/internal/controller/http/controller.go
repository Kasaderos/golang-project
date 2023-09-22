package controller_http

type Services struct {
	OrderCreateService
	OrderInfoService
	OrderPayService
	StockInfoService
	OrderCancelService
}

type Controller struct {
	Services
}

func NewController(us Services) *Controller {
	return &Controller{
		Services: us,
	}
}
