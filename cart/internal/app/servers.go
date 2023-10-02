package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	api "route256/cart/internal/api/carts"
	"route256/cart/internal/middleware/auth"
	desc "route256/cart/pkg/api/carts/v1"
	"route256/cart/pkg/middleware/logging"
	"route256/cart/pkg/middleware/panic"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func initGRPCServer(services *api.Deps) (*grpc.Server, net.Listener, error) {
	lis, err := net.Listen("tcp", os.Getenv("GRPC_ADDR"))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			panic.Interceptor,
			logging.Interceptor,
			auth.Interceptor,
		),
	)

	reflection.Register(grpcServer)

	controller := api.NewServer(services)

	desc.RegisterCartsServer(grpcServer, controller)

	return grpcServer, lis, nil
}

func initGRPCGateway(ctx context.Context, lis net.Listener) (*http.Server, error) {
	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(auth.HeaderMatcher),
	)

	if err := desc.RegisterCartsHandlerFromEndpoint(
		ctx,
		mux,
		lis.Addr().String(),
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
	); err != nil {
		return nil, fmt.Errorf("register carts handler: %v", err)
	}

	return &http.Server{
		Addr:    os.Getenv("GRPC_GW_ADDR"),
		Handler: logging.WithHTTPLoggingMiddleware(mux),
	}, nil
}
