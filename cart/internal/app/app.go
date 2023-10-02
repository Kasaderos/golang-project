package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	mock_repo "route256/cart/internal/repository/mock"
	"syscall"

	"google.golang.org/grpc"
)

func Run() error {
	// Init global context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Client connections
	lomsConn, productsConn, err := initClientConnections(ctx)
	if err != nil {
		return err
	}
	defer closeClientConnections(lomsConn, productsConn)

	// Clients
	lomsClient, productClient := initClients(lomsConn, productsConn)

	// Repository
	cartRepo := mock_repo.NewCartRepostiory()

	// Services
	services := initServices(lomsClient, productClient, cartRepo)

	// Controller
	grpcServer, lis, err := initGRPCServer(services)
	if err != nil {
		return err
	}

	grpcGWServer, err := initGRPCGateway(ctx, lis)
	if err != nil {
		return err
	}

	go initGracefullShutdown(ctx, cancel, grpcServer, grpcGWServer)

	go startGRPCServer(ctx, grpcServer, lis)

	return startGRPCGateway(grpcGWServer)
}

func initGracefullShutdown(
	ctx context.Context,
	cancelGlobalCtx context.CancelFunc,
	grpcServer *grpc.Server,
	grpcGWServer *http.Server,
) {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	<-signalCh

	cancelGlobalCtx()

	grpcServer.GracefulStop()

	if err := grpcGWServer.Shutdown(ctx); err != nil {
		log.Println(err)
	}
}
