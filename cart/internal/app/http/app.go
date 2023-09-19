package http

import (
	"log"
	"net/http"
	"os"
	controller_http "route256/cart/internal/controller/http"
	repository "route256/cart/internal/repository/mock"
	"route256/cart/internal/services/loms"
	"route256/cart/internal/services/product"
	"route256/cart/internal/usecase/cart"
)

func Run() error {
	// Repository
	cartRepo := repository.NewCartRepostiory()

	// Services
	lomsService := loms.NewLOMSService(os.Getenv("LOMS_SERVICE_URL"))
	productService := product.NewProductService(os.Getenv("PRODUCT_SERVICE_URL"))

	// Usecase
	cartUsecase := cart.NewCartUsecase(cart.Deps{
		CartRepository:  cartRepo,
		LOMSService:     lomsService,
		ProductInformer: productService,
	})

	// Controller
	controller := controller_http.NewController(controller_http.Usecases{
		CartService: cartUsecase,
	})

	// Router layer
	router := controller.NewRouter()

	// Run service
	addr := os.Getenv("ADDR")
	log.Printf("cart server is listening at %s", addr)

	return http.ListenAndServe(addr, router)
}
