package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	result1 := make(chan int)
	result2 := make(chan int)

	var sum int // SHARED

	go func() {
		defer wg.Done()
		threadName := "Thread 1"
		result := 0
		for i := 0; i < 5; i++ {
			sum = calculate(threadName, sum, 1, 2)
			result = sum
		}
		result1 <- result
	}()

	go func() {
		defer wg.Done()
		threadName := "Thread 2"
		result := 0
		for i := 0; i < 5; i++ {
			sum = calculate(threadName, sum, 3, 4)
			result = sum
		}
		result2 <- result
	}()

	go func() {
		wg.Wait()
		close(result1)
		close(result2)
	}()

	finalSum := <-result1 + <-result2
	fmt.Println("Result (thread 1 and 2):", finalSum)
}

func calculate(threadName string, previousResult, a, b int) int {
	x := previousResult + a + b
	fmt.Printf("%s: %d\n", threadName, x)
	time.Sleep(5 * time.Millisecond) // Sleep for 0.5 seconds
	return x
}
