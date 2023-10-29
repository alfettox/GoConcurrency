package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	result1 := make(chan int)
	result2 := make(chan int)

	go func() {
		defer wg.Done()
		threadName := "Thread 1"
		result := 0
		for i := 0; i < 5; i++ {
			result = calculate(threadName, result, 1, 2)
		}
		result1 <- result
	}()

	go func() {
		defer wg.Done()
		threadName := "Thread 2"
		result := 0
		for i := 0; i < 5; i++ {
			result = calculate(threadName, result, 3, 4)
		}
		result2 <- result
	}()

	go func() {
		wg.Wait()
		close(result1)
		close(result2)
	}()

	sum := <-result1 + <-result2
	fmt.Println("Result (thread 1 and 2):", sum)
}

func calculate(threadName string, previousResult, a, b int) int {
	x := previousResult + a + b
	fmt.Printf("%s: %d\n", threadName, x)
	return x
}
