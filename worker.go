package main

import (
	"context"
	"fmt"
	"time"
)

// Worker simulates a background task that does work periodically
type Worker struct {
	interval time.Duration
	done     chan struct{}
}

func NewWorker(interval time.Duration) *Worker {
	return &Worker{
		interval: interval,
		done:     make(chan struct{}),
	}
}

// Start begins the worker's periodic task
func (w *Worker) Start(ctx context.Context) error {
	ticker := time.NewTicker(w.interval)
	defer ticker.Stop()

	count := 0
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			count++
			fmt.Printf("worker tick %d at %s\n", count, time.Now().Format("15:04:05"))
		case <-w.done:
			fmt.Println("worker stopped")
			return nil
		}
	}
}

// Stop gracefully stops the worker
func (w *Worker) Stop() {
	close(w.done)
}
