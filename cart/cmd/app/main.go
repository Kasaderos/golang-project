package main

import (
	"log"
	grpc_app "route256/cart/internal/app/grpc"
)

func main() {
	log.Fatal(grpc_app.Run())
}
