package http

import (
	"log"
	"net/http"
	"os"
	"route256/cart/internal/clients/loms"
	"route256/cart/internal/clients/product"
	controller_http "route256/cart/internal/controller/http"
	mock_repository "route256/cart/internal/repository/mock"
	"route256/cart/internal/services/cart"
)

func Run() error {
	// Repository
	cartRepo := mock_repository.NewCartRepostiory()

	// Services
	lomsClient := loms.NewLOMSService(os.Getenv("LOMS_SERVICE_URL"))
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
		ItemDeleter:   cartRepo,
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
