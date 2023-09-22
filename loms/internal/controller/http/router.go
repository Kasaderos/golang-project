package controller_http

import (
	"net/http"
)

func (c *Controller) NewRouter() http.Handler {
	mux := http.NewServeMux()

	handler := MiddlewareRecovery(mux)

	mux.HandleFunc("/order/create", c.OrderCreateHandler)
	mux.HandleFunc("/order/info", c.OrderInfoHandler)
	mux.HandleFunc("/order/pay", c.OrderPayHandler)
	mux.HandleFunc("/order/cancel", c.OrderCancelHandler)
	mux.HandleFunc("/stock/info", c.StockInfoHandler)

	return handler
}
