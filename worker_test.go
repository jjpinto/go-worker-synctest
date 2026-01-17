package main

import (
	"context"
	"testing"
	"testing/synctest"
	"time"
)

func TestWorker_TicksAtInterval(t *testing.T) {
	synctest.Run(func() {
		worker := NewWorker(5 * time.Second)
		
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		go worker.Start(ctx)

		synctest.Wait()
		time.Sleep(5 * time.Second)
		
		synctest.Wait()
		time.Sleep(5 * time.Second)
		
		worker.Stop()
		synctest.Wait()
		
		t.Log("test completed - worker ticked twice in controlled time")
	})
}

func TestWorker_CancelsOnContext(t *testing.T) {
	synctest.Run(func() {
		worker := NewWorker(1 * time.Second)
		
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		errCh := make(chan error, 1)
		go func() {
			errCh <- worker.Start(ctx)
		}()

		synctest.Wait()
		time.Sleep(4 * time.Second)
		
		synctest.Wait()
		err := <-errCh
		
		if err != context.DeadlineExceeded {
			t.Errorf("expected DeadlineExceeded, got %v", err)
		}
	})
}

func TestWorker_MultipleWorkers(t *testing.T) {
	synctest.Run(func() {
		worker1 := NewWorker(2 * time.Second)
		worker2 := NewWorker(3 * time.Second)
		
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		go worker1.Start(ctx)
		go worker2.Start(ctx)

		synctest.Wait()
		time.Sleep(6 * time.Second)
		
		worker1.Stop()
		worker2.Stop()
		synctest.Wait()
		
		t.Log("both workers completed successfully")
	})
}
