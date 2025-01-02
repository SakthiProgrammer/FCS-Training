package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	waitGroupLimit := 10

	wg.Add(waitGroupLimit)

	for i := 0; i < waitGroupLimit; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d started\n", i)
			time.Sleep(time.Second * time.Duration(i)) // Simulate work
			fmt.Printf("Goroutine %d finished\n", i)
		}(i)
	}

	wg.Wait()

	fmt.Println("Completed")
}
