package app

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"google.golang.org/grpc"

	"route256/loms/internal/repository/postgres"
)

type App struct {
	wg sync.WaitGroup
}

func (app *App) Run() error {
	// Init global context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dbpool, err := getDBConnPool(ctx)
	if err != nil {
		return err
	}
	defer dbpool.Close()

	// Repository
	ordersRepo := postgres.NewOrdersRepostiory(dbpool)
	stocksRepo := postgres.NewStocksRepostiory(dbpool)

	// Services
	services := initServices(ordersRepo, stocksRepo)

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

	go app.StartGRPCServer(ctx, grpcServer, lis)

	startGRPCGateway(grpcGWServer)

	app.Wait()

	return nil
}

func (app *App) StartGRPCServer(ctx context.Context, grpcServer *grpc.Server, lis net.Listener) {
	app.wg.Add(1)
	defer app.wg.Done()

	log.Printf("server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
	}
}

func (app *App) Wait() {
	app.wg.Wait()
}

func startGRPCGateway(grcpGateway *http.Server) {
	log.Printf("Serving gRPC-Gateway on %s\n", grcpGateway.Addr)
	if err := grcpGateway.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Printf("failed to serve: %v", err)
		}
	}
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

	log.Println("Service gracefull shutdowning")

	cancelGlobalCtx()

	grpcServer.GracefulStop()

	if err := grpcGWServer.Shutdown(ctx); err != nil {
		log.Println(err)
	}
}
