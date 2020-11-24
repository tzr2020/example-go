package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go work(ctx)

	time.Sleep(3 * time.Second)
	fmt.Println("stop all gorutines.")
	cancel()
	time.Sleep(3 * time.Second)
}

func work(ctx context.Context) {
	go job(ctx)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("work got the stop channel.")
			return
		default:
			fmt.Println("work: do somrthing...")
			time.Sleep(1 * time.Second)
		}
	}
}

func job(ctx context.Context) {
	go task(ctx)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("job got the stop channel.")
			return
		default:
			fmt.Println("job: do somrthing...")
			time.Sleep(1 * time.Second)
		}
	}
}

func task(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("task got the stop channel.")
			return
		default:
			fmt.Println("task: do somrthing...")
			time.Sleep(1 * time.Second)
		}
	}
}
