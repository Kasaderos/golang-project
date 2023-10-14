package workerpool

import (
	"context"
	"log"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

type WorkerPool struct {
	errGroup *errgroup.Group
	sem      *semaphore.Weighted
	ctx      context.Context
}

func New(ctx context.Context, maxWorkers int) *WorkerPool {
	// in this case the new ctx is derived
	errGroup, ctx := errgroup.WithContext(ctx)

	// errGroup.SetLimit is used buffered chan
	// instead this we use semaphore
	sem := semaphore.NewWeighted(int64(maxWorkers))

	return &WorkerPool{
		errGroup: errGroup,
		ctx:      ctx,
		sem:      sem,
	}
}

func (w *WorkerPool) Run(f func() error) {
	// don't run if ctx is done
	select {
	case <-w.ctx.Done():
		return
	default:
	}

	// if we have an error it stops, because ctx is derived
	// if everything is ok we'll be blocked until
	// some worker releases
	if err := w.sem.Acquire(w.ctx, 1); err != nil {
		log.Println("workerpool:", err)
		return
	}

	// RUN
	w.errGroup.Go(func() error {
		defer w.sem.Release(1)
		return f()
	})
}

func (w *WorkerPool) Wait() error {
	return w.errGroup.Wait()
}
