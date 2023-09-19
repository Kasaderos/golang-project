package http

import (
	"log"
	"net/http"
	"os"
	controller_http "route256/loms/internal/controller/http"
	repository "route256/loms/internal/repository/mock"
	oms "route256/loms/internal/usecase/OMS"
	wms "route256/loms/internal/usecase/WMS"
)

func Run() error {
	// Repository
	omsRepo := repository.NewOMSRepostiory()
	wmsRepo := repository.NewStocksRepostiory()

	// Usecase
	omsUsecase := oms.NewOMSUsecase(oms.Deps{
		WMSRepository: wmsRepo,
		OMSRepository: omsRepo,
	})
	wmsUsecase := wms.NewWMSUsecase(wms.Deps{
		WMSRepository: wmsRepo,
	})

	// Controller
	controller := controller_http.NewController(controller_http.Usecases{
		OMSUsecase: omsUsecase,
		WMSUsecase: wmsUsecase,
	})

	// Router layer
	router := controller.NewRouter()

	// Run service
	addr := os.Getenv("ADDR")
	log.Printf("loms server is listening at %s", addr)
	return http.ListenAndServe(addr, router)
}
