package main

import (
	"log"
	http_app "route256/loms/internal/app/http"
)

func main() {
	log.Fatal(http_app.Run())
}
