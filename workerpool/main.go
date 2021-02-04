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
	var workerNum = 5
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < workerNum; i++ {
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
	taskNum := 30
	for i := 0; i < taskNum; i++ {
		channel <- i
	}
	close(channel)
	wg.Wait()

	elapsed := time.Since(startTime)

	fmt.Printf("\nworkerpool took %s to finish %d jobs with %d workers.\n", elapsed, taskNum, workerNum)
}

func main() {
	workerpool()
}
