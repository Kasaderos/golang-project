package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"route256/cart/internal/clients/loms"
	"route256/cart/internal/clients/product"
	products_grpc "route256/cart/pkg/api/products/v1"
	loms_grpc "route256/loms/pkg/api/loms/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func initConnections(ctx context.Context) (loms *grpc.ClientConn, products *grpc.ClientConn, err error) {
	// Init client connections
	lomsConn, err := grpc.DialContext(
		ctx,
		os.Getenv("LOMS_SERVICE_URL"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to LOMS server: %v", err)
	}

	productsConn, err := grpc.DialContext(
		ctx,
		os.Getenv("PRODUCT_SERVICE_URL"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to Products server: %v", err)
	}

	return lomsConn, productsConn, nil
}

func closeConnections(loms *grpc.ClientConn, products *grpc.ClientConn) {
	if err := loms.Close(); err != nil {
		log.Println(err)
	}
	if err := products.Close(); err != nil {
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
