package http

import (
	"log"
	"net"
	"os"
	api "route256/loms/internal/api/loms"
	mock_repository "route256/loms/internal/repository/mock"
	"route256/loms/internal/services/order"
	"route256/loms/internal/services/stock"

	desc "route256/loms/pkg/api/loms/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Run() error {
	// Repository
	omsRepo := mock_repository.NewOMSRepostiory()
	wmsRepo := mock_repository.NewStocksRepostiory()

	// Usecase
	orderCreateService := order.NewCreateService(order.CreateDeps{
		OrderCreator:   omsRepo,
		StocksReserver: wmsRepo,
	})
	orderInfoService := order.NewGetInfoService(omsRepo)
	orderPayService := order.NewPayService(order.PayDeps{
		OrderProvider:     omsRepo,
		ReserveRemover:    wmsRepo,
		OrderStatusSetter: omsRepo,
	})
	orderCancelService := order.NewCancelService(order.CancelDeps{
		OrderProvider:     omsRepo,
		ReserveCanceller:  wmsRepo,
		OrderStatusSetter: omsRepo,
	})
	stocksInfoService := stock.NewStocksService(wmsRepo)

	// GRPC Server
	lis, err := net.Listen("tcp", os.Getenv("ADDR"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	controller := api.NewServer(api.Deps{
		OrderCreateService: orderCreateService,
		OrderInfoService:   orderInfoService,
		OrderPayService:    orderPayService,
		OrderCancelService: orderCancelService,
		StockInfoService:   stocksInfoService,
	})

	desc.RegisterLOMSServer(grpcServer, controller)

	log.Printf("server listening at %v", lis.Addr())

	return grpcServer.Serve(lis)
}
