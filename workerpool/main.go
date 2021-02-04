package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int) {
	fmt.Printf("%d ", i)
}

func workerpool() {
	fmt.Println("============ workerpool ============")
	startTime := time.Now()
	channel := make(chan int, 10)
	var wg sync.WaitGroup
	// 5个worker
	var workerPoolSize = 5
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < workerPoolSize; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for a := range channel {
					process(a)
					time.Sleep(time.Second)
				}
			}()
		}
	}()

	// 30个task
	for i := 0; i < 30; i++ {
		channel <- i
	}
	close(channel)
	wg.Wait()

	elapsed := time.Since(startTime)

	fmt.Printf("\nworkerpool1 took %s to finish\n", elapsed)
}

func main() {
	workerpool()
}
