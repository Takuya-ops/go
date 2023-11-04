package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func f(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		fmt.Println("goroutine:処理")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go f(ctx, &wg)

	time.Sleep(10 * time.Second)
	cancel()
	wg.Wait()
}
