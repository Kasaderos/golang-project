package app

import (
	"log"
	"os"
	api "route256/cart/internal/api/carts"
	"route256/cart/internal/clients/loms"
	"route256/cart/internal/clients/product"
	"route256/cart/internal/repository/postgres"
	"route256/cart/internal/services/cart"
	"strconv"
)

func initServices(
	lomsClient *loms.Client,
	productClient *product.Client,
	cartRepo *postgres.CartRepository,
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

	listItemService := cart.NewListItemService(
		cartRepo,
		productClient,
		getMaxListItemWorkers(),
	)
	clearService := cart.NewClearService(cartRepo)

	return &api.Deps{
		ItemAddService:    addService,
		CheckoutService:   checkoutService,
		ItemDeleteService: itemDeleteService,
		ListItemService:   listItemService,
		ClearService:      clearService,
	}
}

func getMaxListItemWorkers() int {
	const defaultWorkers = 1
	value := os.Getenv("LIST_ITEM_SERVICE_WORKERS")
	workersNum, err := strconv.Atoi(value)
	if err != nil {
		log.Println("config: list service workers unset, using default 1")
		return defaultWorkers
	}

	return workersNum
}
