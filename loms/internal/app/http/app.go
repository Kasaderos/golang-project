package http

import (
	"log"
	"net/http"
	"os"
	controller_http "route256/loms/internal/controller/http"
	mock_repository "route256/loms/internal/repository/mock"
	"route256/loms/internal/services/order"
	"route256/loms/internal/services/stock"
)

func Run() error {
	// Repository
	omsRepo := mock_repository.NewOMSRepostiory()
	wmsRepo := mock_repository.NewStocksRepostiory()

	// Usecase
	orderCreator := order.NewCreateService(order.CreateDeps{
		OrderCreator:   omsRepo,
		StocksReserver: wmsRepo,
	})
	orderInformer := order.NewGetInfoService(omsRepo)
	orderPayer := order.NewPayService(order.PayDeps{
		OrderProvider:     omsRepo,
		ReserveRemover:    wmsRepo,
		OrderStatusSetter: omsRepo,
	})
	orderCanceller := order.NewCancelService(order.CancelDeps{
		OrderProvider:     omsRepo,
		ReserveCanceller:  wmsRepo,
		OrderStatusSetter: omsRepo,
	})
	stocksInformer := stock.NewStocksService(wmsRepo)

	// Controller
	controller := controller_http.NewController(controller_http.Usecases{
		OrderCreator:   orderCreator,
		OrderInformer:  orderInformer,
		OrderPayer:     orderPayer,
		OrderCanceller: orderCanceller,
		StocksInformer: stocksInformer,
	})

	// Router layer
	router := controller.NewRouter()

	// Run service
	addr := os.Getenv("ADDR")
	log.Printf("loms server is listening at %s", addr)
	return http.ListenAndServe(addr, router)
}
