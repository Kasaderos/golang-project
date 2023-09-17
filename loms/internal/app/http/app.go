package http

import (
	"log"
	"net/http"
	"os"
	controller_http "route256/loms/internal/controller/http"
	repository "route256/loms/internal/repository/mock"

	// wms "route256/loms/internal/services/WMS"
	oms "route256/loms/internal/usecase/OMS"
)

type App struct {
	config Config
}

func (a App) Run() {
	// Repository layer
	omsRepo := repository.NewOMSRepostiory( /* ... */ )

	// Other external services (adapters) layer
	wmsRepo := repository.NewStocksRepostiory()

	// Usecase layer
	omsUsecase := oms.NewOMSUsecase(oms.Deps{
		WMSRepository: wmsRepo, // todo
		OMSRepository: omsRepo,
	})

	// Delivery || Gateway || Transport || Controller layer
	controller := controller_http.NewController(controller_http.Usecases{
		OrderManagementSystem: omsUsecase,
	})

	// Router layer
	router := controller.NewRouter()

	// Middleware layer
	// router = middleware.WithHTTPRecoverMiddleware(router)

	// Run service
	addr := os.Getenv("ADDR")
	log.Printf("server is listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
