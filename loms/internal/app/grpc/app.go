package http

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"route256/cart/pkg/middleware/logging"
	api "route256/loms/internal/api/loms"
	mock_repository "route256/loms/internal/repository/mock"
	"route256/loms/internal/services/order"
	"route256/loms/internal/services/stock"
	desc "route256/loms/pkg/api/loms/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	lis, err := net.Listen("tcp", os.Getenv("GRPC_ADDR"))
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

	go func() {
		if err = grpcServer.Serve(lis); err != nil { // запускаем grpc сервер
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Создаем коннект с grpc сервером
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		lis.Addr().String(),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	mux := runtime.NewServeMux()

	err = desc.RegisterLOMSHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    os.Getenv("GRPC_GW_ADDR"),
		Handler: logging.WithHTTPLoggingMiddleware(mux),
	}

	log.Printf("Serving gRPC-Gateway on %s\n", gwServer.Addr)
	return gwServer.ListenAndServe()
}
