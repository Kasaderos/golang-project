package http

import (
	"log"
	"net"
	"os"
	"route256/cart/internal/clients/loms"
	"route256/cart/internal/clients/product"
	mock_repository "route256/cart/internal/repository/mock"
	"route256/cart/internal/services/cart"
	products_grpc "route256/cart/pkg/api/products/v1"
	loms_grpc "route256/loms/pkg/api/loms/v1"

	api "route256/cart/internal/api/carts"
	desc "route256/cart/pkg/api/carts/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	productsConn, err := grpc.Dial(os.Getenv("PRODUCTS_SERVICE_URL"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to Products server: %v", err)
	}
	defer productsConn.Close()

	grpcLOMSClient := loms_grpc.NewLOMSClient(lomsConn)
	grpcProductsClient := products_grpc.NewProductsClient(productsConn)

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
	lis, err := net.Listen("tcp", os.Getenv("ADDR"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// reflection.Register(grpcServer)

	controller := api.NewServer(api.Deps{
		ItemAddService:    addService,
		CheckoutService:   checkoutService,
		ItemDeleteService: itemDeleteService,
		ListItemService:   listItemService,
		ClearService:      clearService,
	})

	desc.RegisterCartsServer(grpcServer, controller)

	log.Printf("server listening at %v", lis.Addr())

	return grpcServer.Serve(lis)
}
