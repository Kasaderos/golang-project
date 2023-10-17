package workerpool

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

func TestWorkerPool(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	const (
		maxWorkers = 20
		jobsNum    = 10
	)

	sum := int64(0)

	wp, _ := New(ctx, maxWorkers)
	for i := 0; i < jobsNum; i++ {
		wp.Run(func() error {
			atomic.AddInt64(&sum, 1)
			time.Sleep(10 * time.Millisecond)
			return nil
		})
	}

	wp.Wait()

	if sum != jobsNum {
		t.Errorf("expected %d, actual %d", jobsNum, sum)
	}
}

func TestWorkerPoolCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	const (
		maxWorkers = 20
		jobsNum    = 10
	)

	sum := int64(0)

	wp, _ := New(ctx, maxWorkers)
	for i := 0; i < jobsNum; i++ {
		cancel()
		wp.Run(func() error {
			atomic.AddInt64(&sum, 1)
			time.Sleep(10 * time.Millisecond)
			return nil
		})
	}

	wp.Wait()

	if sum != 0 {
		t.Errorf("expected 0, actual %d", sum)
	}
}
