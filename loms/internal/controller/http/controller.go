package controller_http

type Usecases struct {
	OrderCreator
	OrderInformer
	OrderPayer
	StocksInformer
	OrderCanceller
}

type Controller struct {
	Usecases
}

func NewController(us Usecases) *Controller {
	return &Controller{
		Usecases: us,
	}
}
