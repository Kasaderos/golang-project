package app

import (
	api "route256/loms/internal/api/loms"
	"route256/loms/internal/repository/postgres"
	"route256/loms/internal/services/notification"
	"route256/loms/internal/services/order"
	"route256/loms/internal/services/stock"

	"github.com/Shopify/sarama"
)

func initServices(
	ordersRepo *postgres.OrdersRepository,
	stocksRepo *postgres.StocksRepository,
	producer sarama.SyncProducer,
) *api.Deps {
	notificationService := notification.NewService(producer)

	orderCreateService := order.NewCreateService(order.CreateDeps{
		OrderCreator:      ordersRepo,
		StocksReserver:    stocksRepo,
		OrderStatusSetter: ordersRepo,
		StatusNotifier:    notificationService,
	})

	orderPayService := order.NewPayService(order.PayDeps{
		OrderProvider:     ordersRepo,
		ReserveRemover:    stocksRepo,
		OrderStatusSetter: ordersRepo,
	})

	orderCancelService := order.NewCancelService(order.CancelDeps{
		OrderProvider:     ordersRepo,
		ReserveCanceller:  stocksRepo,
		OrderStatusSetter: ordersRepo,
		StatusNotifier:    notificationService,
	})

	orderInfoService := order.NewGetInfoService(ordersRepo)
	stocksInfoService := stock.NewStocksService(stocksRepo)

	return &api.Deps{
		OrderCreateService: orderCreateService,
		OrderInfoService:   orderInfoService,
		OrderPayService:    orderPayService,
		OrderCancelService: orderCancelService,
		StockInfoService:   stocksInfoService,
	}
}
