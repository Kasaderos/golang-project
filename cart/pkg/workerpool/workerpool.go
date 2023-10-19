package workerpool

import (
	"context"
	"runtime"
	"sync"
	"time"
)

// WorkerTTL is timeout for obtaining the next job
// After this timeout worker will stop
const WorkerTTL = 100 * time.Microsecond

type signal struct{}
type worker struct{}

type job struct {
	accepted chan signal
	f        func() error
}

// WorkerPool is some of kind of errgroup
type WorkerPool struct {
	wg            sync.WaitGroup
	err           error
	errOnce       sync.Once
	jobs          chan job
	ctx           context.Context
	cancel        context.CancelFunc
	activeWorkers chan worker
}

// New is constructor of WorkerPool.
// size is the maximum workers(goroutines).
func New(ctx context.Context, size int) (*WorkerPool, context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	return &WorkerPool{
		jobs:          make(chan job, 1),
		ctx:           ctx,
		cancel:        cancel,
		activeWorkers: make(chan worker, size),
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

	jb := job{
		accepted: make(chan signal, 1),
		f:        f,
	}

	select {
	case wp.jobs <- jb:
		// try to activate worker
		select {
		case <-jb.accepted:
			// some worker accepted
		case wp.activeWorkers <- worker{}:
			// worker size isn't reached
			// we can run a new worker
			// there are two scenarios:
			// scenario 1:
			//  it started and it got the job
			// scenario 2:
			//  it started, but another worker got job
			wp.wg.Add(1)
			go wp.worker()
		}
	case <-wp.ctx.Done():
	}
}

// we assume that the workers don't fall [or calls job() don't panic]
// we can add recovery, but it will be too difficult ...
func (wp *WorkerPool) worker() {
	defer wp.wg.Done()
	// instead time.After I decided to use flag
	// try to get some job after
	retry := true
	for {
		// check context
		select {
		case <-wp.ctx.Done():
			return
		default:
		}

		// try to get some job
		select {
		case job, hasJob := <-wp.jobs:
			if !hasJob {
				return
			}
			// send signal that we accepted the job
			job.accepted <- signal{}

			if err := job.f(); err != nil {
				wp.errOnce.Do(func() {
					wp.err = err
					// it's bad idea to cancel itself,
					// but in this case we can do this
					// we just want to say that 'guys, stop working'
					wp.cancel()
				})
			}
			runtime.Gosched()
			retry = true
		default:
			// retry get a job after WorkerTTL
			if retry {
				time.Sleep(WorkerTTL)
				retry = false
				continue
			}

			// we didn't get any job after retry
			// so let's quit and stop goroutine
			<-wp.activeWorkers
			return
		}
	}
}

// Wait waits all active goroutines.
func (wp *WorkerPool) Wait() error {
	close(wp.jobs)
	wp.wg.Wait()

	return wp.err
}
