package http

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	api "route256/cart/internal/api/carts"
	"route256/cart/internal/clients/loms"
	"route256/cart/internal/clients/product"
	"route256/cart/internal/middleware/auth"
	mock_repo "route256/cart/internal/repository/mock"
	"route256/cart/internal/services/cart"
	desc "route256/cart/pkg/api/carts/v1"
	products_grpc "route256/cart/pkg/api/products/v1"
	"route256/cart/pkg/middleware/logging"
	"route256/cart/pkg/middleware/panic"
	loms_grpc "route256/loms/pkg/api/loms/v1"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func Run() error {
	// Init global context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Client connections
	lomsConn, productsConn, err := initConnections(ctx)
	if err != nil {
		return err
	}
	defer closeConnections(lomsConn, productsConn)

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

func initConnections(ctx context.Context) (loms *grpc.ClientConn, products *grpc.ClientConn, err error) {
	// Init client connections
	lomsConn, err := grpc.DialContext(
		ctx,
		os.Getenv("LOMS_SERVICE_URL"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to LOMS server: %v", err)
	}

	productsConn, err := grpc.DialContext(
		ctx,
		os.Getenv("PRODUCT_SERVICE_URL"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to Products server: %v", err)
	}

	return lomsConn, productsConn, nil
}

func closeConnections(loms *grpc.ClientConn, products *grpc.ClientConn) {
	if err := loms.Close(); err != nil {
		log.Println(err)
	}
	if err := products.Close(); err != nil {
		log.Println(err)
	}
}

func initClients(
	lomsConn grpc.ClientConnInterface,
	productsConn grpc.ClientConnInterface,
) (*loms.Client, *product.Client) {
	grpcLOMSClient := loms_grpc.NewLOMSClient(lomsConn)
	grpcProductsClient := products_grpc.NewProductServiceClient(productsConn)

	lomsClient := loms.NewClient(grpcLOMSClient)
	productClient := product.NewClient(grpcProductsClient)

	return lomsClient, productClient
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

func initServices(
	lomsClient *loms.Client,
	productClient *product.Client,
	cartRepo *mock_repo.CartRepository,
) *api.Deps {
	addService := cart.NewAddService(cart.AddDeps{
		ProductProvider: productClient,
		StockProvider:   lomsClient,
		ItemAdder:       cartRepo,
	})
	checkoutService := cart.NewCheckoutService(cart.CheckoutDeps{
		OrderCreator:  lomsClient,
		ItemsProvider: cartRepo,
		ItemsDeleter:  cartRepo,
	})
	itemDeleteService := cart.NewItemDeleteService(cartRepo)
	listItemService := cart.NewListItemService(cartRepo, productClient)
	clearService := cart.NewClearService(cartRepo)

	return &api.Deps{
		ItemAddService:    addService,
		CheckoutService:   checkoutService,
		ItemDeleteService: itemDeleteService,
		ListItemService:   listItemService,
		ClearService:      clearService,
	}
}

func initGRPCServer(services *api.Deps) (*grpc.Server, net.Listener, error) {
	lis, err := net.Listen("tcp", os.Getenv("GRPC_ADDR"))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			panic.Interceptor,
			logging.Interceptor,
			auth.Interceptor,
		),
	)

	reflection.Register(grpcServer)

	controller := api.NewServer(services)

	desc.RegisterCartsServer(grpcServer, controller)

	return grpcServer, lis, nil
}

func startGRPCServer(ctx context.Context, grpcServer *grpc.Server, lis net.Listener) {
	log.Printf("server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
	}
}

func initGRPCGateway(ctx context.Context, lis net.Listener) (*http.Server, error) {
	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(auth.HeaderMatcher),
	)

	if err := desc.RegisterCartsHandlerFromEndpoint(ctx, mux, lis.Addr().String(), []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}); err != nil {
		return nil, fmt.Errorf("register carts handler: %v", err)
	}

	return &http.Server{
		Addr:    os.Getenv("GRPC_GW_ADDR"),
		Handler: logging.WithHTTPLoggingMiddleware(mux),
	}, nil
}

func startGRPCGateway(grcpGateway *http.Server) error {
	log.Printf("Serving gRPC-Gateway on %s\n", grcpGateway.Addr)
	if err := grcpGateway.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
