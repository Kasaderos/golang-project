package main

import (
	"log"
	grpc_app "route256/loms/internal/app/grpc"
)

func main() {
	log.Fatal(grpc_app.Run())
}
