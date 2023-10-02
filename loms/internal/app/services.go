package app

import (
	api "route256/loms/internal/api/loms"
	mock_repo "route256/loms/internal/repository/mock"
	"route256/loms/internal/services/order"
	"route256/loms/internal/services/stock"
)

func initServices(
	ordersRepo *mock_repo.OrdersRepository,
	stocksRepo *mock_repo.StocksRepository,
) *api.Deps {
	orderCreateService := order.NewCreateService(order.CreateDeps{
		OrderCreator:   ordersRepo,
		StocksReserver: stocksRepo,
	})
	orderInfoService := order.NewGetInfoService(ordersRepo)
	orderPayService := order.NewPayService(order.PayDeps{
		OrderProvider:     ordersRepo,
		ReserveRemover:    stocksRepo,
		OrderStatusSetter: ordersRepo,
	})
	orderCancelService := order.NewCancelService(order.CancelDeps{
		OrderProvider:     ordersRepo,
		ReserveCanceller:  stocksRepo,
		OrderStatusSetter: ordersRepo,
	})
	stocksInfoService := stock.NewStocksService(stocksRepo)

	return &api.Deps{
		OrderCreateService: orderCreateService,
		OrderInfoService:   orderInfoService,
		OrderPayService:    orderPayService,
		OrderCancelService: orderCancelService,
		StockInfoService:   stocksInfoService,
	}
}
