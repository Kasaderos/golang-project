package controller_http

import (
	"net/http"
	"route256/cart/pkg/middleware"
	"route256/cart/pkg/router"
)

func (c *Controller) NewRouter() http.Handler {
	mux := router.NewRouter()

	handler := middleware.Recovery(mux)

	mux.HandleFunc(http.MethodPost, "/order/create", c.OrderCreateHandler)
	mux.HandleFunc(http.MethodPost, "/order/info", c.OrderInfoHandler)
	mux.HandleFunc(http.MethodPost, "/order/pay", c.OrderPayHandler)
	mux.HandleFunc(http.MethodPost, "/order/cancel", c.OrderCancelHandler)
	mux.HandleFunc(http.MethodPost, "/stock/info", c.StockInfoHandler)

	return handler
}
