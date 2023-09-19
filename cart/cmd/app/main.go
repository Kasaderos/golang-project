package main

import (
	"log"
	httpapp "route256/cart/internal/app/http"
)

func main() {
	log.Fatal(httpapp.Run())
}
