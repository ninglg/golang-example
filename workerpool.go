package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int) {
	fmt.Println(i)
}

func main() {
	start := time.Now()
	channel := make(chan int, 100)
	var wg sync.WaitGroup
	var workerPoolSize = 100
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < workerPoolSize; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for a := range channel {
					process(a)
				}
			}()
		}
	}()

	for i := 0; i < 1000; i++ {
		channel <- i
	}
	close(channel)
	wg.Wait()

	elasped := time.Since(start)

	fmt.Printf("Took %s\n", elasped)
}
