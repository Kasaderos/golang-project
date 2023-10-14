package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"route256/cart/internal/clients/loms"
	"route256/cart/internal/clients/product"
	products_grpc "route256/cart/pkg/api/products/v1"
	rate "route256/cart/pkg/middleware/rate"
	loms_grpc "route256/loms/pkg/api/loms/v1"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func initClientConnections(
	ctx context.Context,
) (lomsConn *grpc.ClientConn, productsConn *grpc.ClientConn, err error) {
	// Init client connections
	lomsConn, err = grpc.DialContext(
		ctx,
		os.Getenv("LOMS_SERVICE_URL"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to LOMS server: %v", err)
	}

	limiter := rate.New(getProductServiceRPS())
	productsConn, err = grpc.DialContext(
		ctx,
		os.Getenv("PRODUCT_SERVICE_URL"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(limiter.RequestInterceptor),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to Products server: %v", err)
	}

	return lomsConn, productsConn, nil
}

func closeClientConnections(lomsConn *grpc.ClientConn, productsConn *grpc.ClientConn) {
	if err := lomsConn.Close(); err != nil {
		log.Println(err)
	}
	if err := productsConn.Close(); err != nil {
		log.Println(err)
	}
}

func initClients(
	lomsConn grpc.ClientConnInterface,
	productsConn grpc.ClientConnInterface,
) (*loms.Client, *product.Client) {
	grpcLOMSClient := loms_grpc.NewLOMSClient(lomsConn)
	grpcProductsClient := products_grpc.NewProductServiceClient(productsConn)

	lomsClient := loms.NewClient(grpcLOMSClient)
	productClient := product.NewClient(grpcProductsClient)

	return lomsClient, productClient
}

func getProductServiceRPS() int {
	const defaultRPS = 10
	value := os.Getenv("PRODUCT_SERVICE_RPS")
	rps, err := strconv.Atoi(value)
	if err != nil {
		log.Println("config: product service rate limiter unset, using default 10 RPS")
		return defaultRPS
	}
	return rps
}
