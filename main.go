package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	fmt.Println("starting worker demo...")

	worker := NewWorker(2 * time.Second)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	go func() {
		<-sigChan
		fmt.Println("\nreceived interrupt, stopping...")
		cancel()
	}()

	fmt.Println("worker running (press Ctrl+C to stop)...")
	if err := worker.Start(ctx); err != nil {
		fmt.Printf("worker error: %v\n", err)
	}

	fmt.Println("bye!")
}
