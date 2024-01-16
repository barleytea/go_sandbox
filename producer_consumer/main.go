package main

import (
	"context"
	"log"
	"os"
	"runtime/trace"
	"sync"
)

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
	ch := make(chan int, 10)

	// 1 ~ 100 までの整数を channel に送信する
	go produce(100, ch, ctx)

	// channel からデータを受信する
	for i := 0; i < CONSUMERS_COUNT; i++ {
		i := i
		wg.Add(1)
		go consume(i, ch, &wg, ctx)
	}

	wg.Wait()
}

func produce(num int, ch chan int, ctx context.Context) {
	defer trace.StartRegion(ctx, "produce").End()
	var pg sync.WaitGroup
	defer close(ch)
	for i := 1; i <= num; i++ {
		i := i
		pg.Add(1)
		go func() {
			defer pg.Done()
			ch <- i
		}()
	}
	pg.Wait()
}

func consume(idx int, ch chan int, wg *sync.WaitGroup, ctx context.Context) {
	defer trace.StartRegion(ctx, "consume").End()
	defer wg.Done()
	for i := range ch {
		log.Printf("#%d: consumed %d\n", idx, i)
	}
}
