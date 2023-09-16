package main

import (
	"log"
	"net/http"

	"route256/cart/internal/pkg/clients/loms"
	"route256/cart/internal/pkg/clients/product"
	hitem "route256/cart/internal/pkg/handlers/item"
	sitem "route256/cart/internal/pkg/services/item"
)

func main() {
	lomsClient, err := loms.New("loms client", "http://loms:8080")
	if err != nil {
		log.Fatal(err)
	}
	productClient, err := product.New("product client", "http://route256.pavl.uk:8080")
	if err != nil {
		log.Fatal(err)
	}

	itemAddHandler := hitem.NewItemAddHandler(sitem.NewAddService(lomsClient, productClient))
	http.HandleFunc("/item/add", itemAddHandler.Handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
