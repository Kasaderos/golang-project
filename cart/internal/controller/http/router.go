package http

import (
	"net/http"
)

func (c *Controller) NewRouter() http.Handler {
	mux := http.NewServeMux()

	// mux.HandleFunc("/order/create", c.CreateOrderHandler)
	// mux.HandleFunc("/order/info", c.OrderInfoHandler)
	// mux.HandleFunc("/order/pay", c.OrderPayHandler)
	// mux.HandleFunc("/order/cancel", c.CancelOrderHandler)
	// mux.HandleFunc("/stock/info", c.StockInfoHandler)

	return mux
}
