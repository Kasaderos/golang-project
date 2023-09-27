package http

import (
	"log"
	"net/http"
	"os"
	"route256/cart/internal/clients/http/product"
	"route256/cart/internal/clients/loms"
	controller_http "route256/cart/internal/controller/http"
	mock_repository "route256/cart/internal/repository/mock"
	"route256/cart/internal/services/cart"
	loms_grpc "route256/cart/pkg/api/loms/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Run() error {
	// Repository
	cartRepo := mock_repository.NewCartRepostiory()

	// clients
	lomsConn, err := grpc.Dial(os.Getenv("LOMS_GRPC_URL"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to LOMS: %v", err)
	}
	defer lomsConn.Close()

	grpcLOMSClient := loms_grpc.NewLOMSClient(lomsConn)

	lomsClient := loms.NewClient(grpcLOMSClient)
	productClient := product.NewProductService(os.Getenv("PRODUCT_SERVICE_URL"))

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

	// Controller
	controller := controller_http.NewController(controller_http.Services{
		ItemAddService:    addService,
		CheckoutService:   checkoutService,
		ItemDeleteService: itemDeleteService,
		ListItemService:   listItemService,
		ClearService:      clearService,
	})

	// Router layer
	router := controller.NewRouter()

	// Run service
	addr := os.Getenv("ADDR")
	log.Printf("cart server is listening at %s", addr)

	return http.ListenAndServe(addr, router)
}
