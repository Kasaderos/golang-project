package controller_http

import (
	"net/http"
	"route256/cart/pkg/middleware"
	"route256/cart/pkg/router"
)

func (c *Controller) NewRouter() http.Handler {
	mux := router.NewRouter()

	handler := middleware.Recovery(mux)

	mux.HandleFunc(http.MethodPost, "/cart/item/add", c.ItemAddHandler)
	mux.HandleFunc(http.MethodPost, "/cart/item/delete", c.ItemDeleteHandler)
	mux.HandleFunc(http.MethodPost, "/cart/list", c.ListHandler)
	mux.HandleFunc(http.MethodPost, "/cart/clear", c.ClearHandler)
	mux.HandleFunc(http.MethodPost, "/cart/checkout", c.CheckoutHandler)

	return handler
}
