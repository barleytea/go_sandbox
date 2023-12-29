package main

import (
	"context"
	"log"
	"os"
	"runtime/trace"
	"sync"

	"golang.org/x/tools/go/analysis/passes/defers"
)

const MAX_PRODUCERS_COUNT_PER_CHANNEL = 10
const CONSUMERS_COUNT = 10

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatal(err)
	}
	defer trace.Stop()

	process()
}

func process() {
	ctx, task := trace.NewTask(context.Background(), "process")
	defer task.End()

	var wg sync.WaitGroup
	ch1 := make(chan int, MAX_PRODUCERS_COUNT_PER_CHANNEL)
	ch2 := make(chan int, MAX_PRODUCERS_COUNT_PER_CHANNEL)

	// 1 ~ 100 までの整数を channel に送信する
	go produce(100, ch1, ctx)

	// channel からデータを受信する
	for i := 0; i < CONSUMERS_COUNT; i++ {
		i := i
		wg.Add(1)
		go consume(i, ch1, &wg, ctx)
	}

	wg.Wait()
}

func produce(num int, ch1 chan int, ch2 chan int, ctx context.Context) {
	defer trace.StartRegion(ctx, "produce").End()
	var pg sync.WaitGroup
	defer close(ch1)
	defer close(ch2)
	for i := 0; i < num; i++ {
		i := i
		pg.Add(1)
		go func() {
			defer pg.Done()
			select {
			case ch1 <- i:
				// do nothing
			case ch2 <- i:
				// do nothing
			default:
				// do nothing
			}
			log.Printf("procuded %d\n", i)
		}()
	}
	pg.Wait()
}

func consume(idx int, ch1 chan int, ch2 chan int, wg *sync.WaitGroup, ctx context.Context) {
	defer trace.StartRegion(ctx, "consume").End()
	defer wg.Done()
	for i := range ch {
		log.Printf("consumer#%d consumed %d\n", idx, i)
	}
}
