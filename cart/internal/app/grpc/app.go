package http

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	api "route256/cart/internal/api/carts"
	"route256/cart/internal/clients/loms"
	"route256/cart/internal/clients/product"
	"route256/cart/internal/middleware/auth"
	mock_repository "route256/cart/internal/repository/mock"
	"route256/cart/internal/services/cart"
	desc "route256/cart/pkg/api/carts/v1"
	products_grpc "route256/cart/pkg/api/products/v1"
	"route256/cart/pkg/middleware/logging"
	"route256/cart/pkg/middleware/panic"
	loms_grpc "route256/loms/pkg/api/loms/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func Run() error {
	// Repository
	cartRepo := mock_repository.NewCartRepostiory()

	// clients
	lomsConn, err := grpc.Dial(os.Getenv("LOMS_SERVICE_URL"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to LOMS server: %v", err)
	}
	defer lomsConn.Close()

	productsConn, err := grpc.Dial(os.Getenv("PRODUCT_SERVICE_URL"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to Products server: %v", err)
	}
	defer productsConn.Close()

	grpcLOMSClient := loms_grpc.NewLOMSClient(lomsConn)
	grpcProductsClient := products_grpc.NewProductServiceClient(productsConn)

	lomsClient := loms.NewClient(grpcLOMSClient)
	productClient := product.NewClient(grpcProductsClient)

	// Usecase
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

	// GRPC Server
	lis, err := net.Listen("tcp", os.Getenv("GRPC_ADDR"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor( // Unary интерсепторы (aka middleware)
			panic.Interceptor,
			logging.Interceptor,
			auth.Interceptor,
		),
	)

	reflection.Register(grpcServer)

	controller := api.NewServer(api.Deps{
		ItemAddService:    addService,
		CheckoutService:   checkoutService,
		ItemDeleteService: itemDeleteService,
		ListItemService:   listItemService,
		ClearService:      clearService,
	})

	desc.RegisterCartsServer(grpcServer, controller)

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

	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(auth.HeaderMatcher),
	)

	err = desc.RegisterCartsHandler(context.Background(), mux, conn)
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
