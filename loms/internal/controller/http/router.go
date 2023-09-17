package controller_http

import (
	"net/http"
)

func (c *Controller) NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/order/create", c.CreateOrderHandler)

	return mux
}
