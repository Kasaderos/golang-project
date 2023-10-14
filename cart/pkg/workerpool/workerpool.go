package workerpool

import (
	"context"
	"runtime"
	"sync"
)

// WorkerPool is some of kind of errgroup
type WorkerPool struct {
	wg      sync.WaitGroup
	err     error
	errOnce sync.Once
	jobs    chan func() error
	ctx     context.Context
	cancel  context.CancelFunc

	size int64
	cur  int64
}

// New is constructor of WorkerPool.
// size is the maximum workers(goroutines).
func New(ctx context.Context, size int) (*WorkerPool, context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	return &WorkerPool{
		jobs:   make(chan func() error, 1),
		ctx:    ctx,
		cancel: cancel,
		size:   int64(size),
	}, ctx
}

// Run calls the given function in a new goroutine.
// It isn't thread safe, so it should be called in the main thread
// You should pass ctx to f somehow to prevent goroutine leak.
func (wp *WorkerPool) Run(f func() error) {
	// prevent the next runs if ctx is cancelled
	select {
	case <-wp.ctx.Done():
		return
	default:
	}

	// check if size is reached
	if wp.cur < wp.size {
		wp.cur++
		wp.wg.Add(1)
		go wp.worker()
	}

	// block if we can push some job or just stop by ctx
	// it can be push the f once randomly, although we have ctx.Done.
	select {
	case wp.jobs <- f:
	case <-wp.ctx.Done():
	}
}

// we assume that the workers don't fall [or calls job() don't panic]
// we can add recovery, but it will be too difficult ...
func (wp *WorkerPool) worker() {
	defer wp.wg.Done()
	for {
		select {
		case <-wp.ctx.Done():
			return
		case job, ok := <-wp.jobs:
			if !ok {
				return
			}
			if err := job(); err != nil {
				wp.errOnce.Do(func() {
					wp.err = err
					// it's bad idea to cancel itself,
					// but in this case we can do this
					// we just want to say that 'guys, stop working'
					wp.cancel()
				})
			}
			runtime.Gosched()
		}
	}
}

// Wait waits all active goroutines.
func (wp *WorkerPool) Wait() error {
	close(wp.jobs)
	wp.wg.Wait()

	return wp.err
}
