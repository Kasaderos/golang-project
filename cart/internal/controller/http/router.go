package controller_http

import (
	"net/http"
)

func (c *Controller) NewRouter() http.Handler {
	mux := http.NewServeMux()

	// mux.HandleFunc("/cart/item/add", c.)
	// mux.HandleFunc("/cart/item/delete", c.OrderInfoHandler)
	// mux.HandleFunc("/cart/list", c.OrderPayHandler)
	// mux.HandleFunc("/cart/clear", c.CancelOrderHandler)
	// mux.HandleFunc("/cart/checkout", c.StockInfoHandler)

	return mux
}
