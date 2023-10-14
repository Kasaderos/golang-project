package rate

import (
	"context"

	"golang.org/x/time/rate"
	"google.golang.org/grpc"
)

// burstSize is the max number of requests per second
const burstSize = 1

type Interceptor struct {
	limiter *rate.Limiter
}

// New is constructor of Interceptor.
// It allows you limiting requests to the grpc client.
// rps is the number of requests per second.
func New(rps int) *Interceptor {
	return &Interceptor{
		limiter: rate.NewLimiter(rate.Limit(rps), burstSize),
	}
}

func (r Interceptor) RequestInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	if err := r.limiter.Wait(ctx); err != nil {
		return err
	}
	return invoker(ctx, method, req, reply, cc, opts...)
}
