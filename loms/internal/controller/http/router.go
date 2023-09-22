package controller_http

import (
	"net/http"
	"route256/cart/pkg/middleware"
)

func (c *Controller) NewRouter() http.Handler {
	mux := http.NewServeMux()

	handler := middleware.Recovery(mux)

	mux.HandleFunc("/order/create", c.OrderCreateHandler)
	mux.HandleFunc("/order/info", c.OrderInfoHandler)
	mux.HandleFunc("/order/pay", c.OrderPayHandler)
	mux.HandleFunc("/order/cancel", c.OrderCancelHandler)
	mux.HandleFunc("/stock/info", c.StockInfoHandler)

	return handler
}
