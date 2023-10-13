package auth

import (
	"context"
	"errors"
	"route256/cart/internal/clients/product"
	servicepb "route256/cart/pkg/api/carts/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const ProductServiceMetaKey = "X-Product-Service-Token"

var ErrProductServiceTokenRequired = errors.New("X-Product-Service-Token required")

// RPC methods that interacts with Product Service
var (
	needProductServiceTokenMethods = map[string]struct{}{
		servicepb.Carts_ItemAdd_FullMethodName: {},
		servicepb.Carts_List_FullMethodName:    {},
	}
)

func Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	if _, yes := needProductServiceTokenMethods[info.FullMethod]; yes {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, ErrProductServiceTokenRequired
		}

		values := md.Get(ProductServiceMetaKey)
		if len(values) < 1 {
			return nil, ErrProductServiceTokenRequired
		}

		ctx = product.WithToken(ctx, values[0])
	}

	return handler(ctx, req)
}
