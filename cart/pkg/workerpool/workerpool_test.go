package workerpool

import (
	"context"
	"io"
	"sync/atomic"
	"testing"
	"time"
)

func TestWorkerPool(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	testCases := []struct {
		MaxWorkers int
		JobsNum    int
	}{
		{
			MaxWorkers: 2,
			JobsNum:    0,
		},
		{
			MaxWorkers: 1,
			JobsNum:    1,
		},
		{
			MaxWorkers: 1,
			JobsNum:    10,
		},
		{
			MaxWorkers: 10,
			JobsNum:    5,
		},
	}

	for _, tt := range testCases {
		sum := int64(0)

		wp, _ := New(ctx, tt.MaxWorkers)
		for i := 0; i < tt.JobsNum; i++ {
			wp.Run(func() error {
				atomic.AddInt64(&sum, 1)
				time.Sleep(10 * time.Millisecond)
				return nil
			})
		}

		_ = wp.Wait()

		if sum != int64(tt.JobsNum) {
			t.Errorf("expected %d, actual %d", tt.JobsNum, sum)
		}
	}
}

func TestWorkerPoolCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	const maxWorkers = 20

	sum := int64(0)

	wp, _ := New(ctx, maxWorkers)
	cancel()
	wp.Run(func() error {
		atomic.AddInt64(&sum, 1)
		time.Sleep(10 * time.Millisecond)
		return nil
	})

	_ = wp.Wait()

	if sum != 0 {
		t.Errorf("expected 0, actual %d", sum)
	}
}

func TestWorkerPoolWithError(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	const (
		maxWorkers  = 20
		expectedSum = 2
	)

	sum := int64(0)

	wp, _ := New(ctx, maxWorkers)
	wp.Run(func() error {
		atomic.AddInt64(&sum, 1)
		time.Sleep(10 * time.Millisecond)
		return nil
	})
	wp.Run(func() error {
		atomic.AddInt64(&sum, 1)
		time.Sleep(10 * time.Millisecond)
		return io.ErrUnexpectedEOF
	})

	err := wp.Wait()
	if err == nil {
		t.Errorf("expected err, got nil")
	}

	if sum != expectedSum {
		t.Errorf("expected %d, actual %d", expectedSum, sum)
	}
}
