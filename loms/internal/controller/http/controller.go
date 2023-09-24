package controller_http

type Services struct {
	OrderCreateService
	OrderInfoService
	OrderPayService
	StockInfoService
	OrderCancelService
}

type Controller struct {
	orderCreateService OrderCreateService
	orderInfoService   OrderInfoService
	orderPayService    OrderPayService
	stockInfoService   StockInfoService
	orderCancelService OrderCancelService
}

func NewController(srvs Services) *Controller {
	return &Controller{
		orderCreateService: srvs.OrderCreateService,
		orderInfoService:   srvs.OrderInfoService,
		orderPayService:    srvs.OrderPayService,
		stockInfoService:   srvs.StockInfoService,
		orderCancelService: srvs.OrderCancelService,
	}
}
