package app

import (
	api "route256/cart/internal/api/carts"
	"route256/cart/internal/clients/loms"
	"route256/cart/internal/clients/product"
	mock_repo "route256/cart/internal/repository/mock"
	"route256/cart/internal/services/cart"
)

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
