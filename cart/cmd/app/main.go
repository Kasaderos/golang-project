package main

import (
	"flag"
	"log"
	"net/http"

	"route256/cart/internal/pkg/clients/loms"
	"route256/cart/internal/pkg/clients/product"
	hitem "route256/cart/internal/pkg/handlers/item"
	sitem "route256/cart/internal/pkg/services/item"
)

func main() {
	opts := newOptions()

	lomsClient, err := loms.New("loms client", opts.lomsAddr)
	if err != nil {
		log.Fatal(err)
	}
	productClient, err := product.New("product client", opts.productAddr)
	if err != nil {
		log.Fatal(err)
	}

	itemAddHandler := hitem.NewItemAddHandler(sitem.NewAddService(lomsClient, productClient))
	http.HandleFunc("/item/add", itemAddHandler.Handle)
	log.Fatal(http.ListenAndServe(opts.addr, nil))
}

type options struct {
	addr        string
	lomsAddr    string
	productAddr string
}

func newOptions() *options {
	const (
		defaultAddr        = ":8080"
		defaultLomsAddr    = "http://loms:8080"
		defaultProductAddr = "http://route256.pavl.uk:8080"
	)

	result := &options{}
	flag.StringVar(&result.addr, "addr", defaultAddr, "server address, default: "+defaultAddr)
	flag.StringVar(&result.lomsAddr, "loms_addr", defaultLomsAddr, "loms address, default: "+defaultLomsAddr)
	flag.StringVar(&result.productAddr, "product_addr", defaultProductAddr, "product-service address, default: "+defaultProductAddr)
	flag.Parse()
	return result
}
