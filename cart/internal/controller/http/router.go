package controller_http

import (
	"net/http"
)

func (c *Controller) NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/cart/item/add", c.ItemAddHandler)
	mux.HandleFunc("/cart/item/delete", c.ItemDeleteHandler)
	mux.HandleFunc("/cart/list", c.ListHandler)
	mux.HandleFunc("/cart/clear", c.ClearHandler)
	mux.HandleFunc("/cart/checkout", c.CheckoutHandler)

	return mux
}
