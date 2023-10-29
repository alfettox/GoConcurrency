package main

import (
	"fmt"
	"sync"
	"time"
	"os"
	"strconv"
)

// TO TEST RACE CONDITIONS: go run -race mutex.go

// MUTEX = Mutual exclusion == Java Lock

const delay = 400*time.Millisecond //milliseconds

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	result1 := make(chan int)
	result2 := make(chan int)

	var sumProtected, sumUnprotected int

	var protRaceCondition bool
	if len(os.Args) > 1 && os.Args[1] == "true" {
		protRaceCondition = true
	} else {
		protRaceCondition = false
	}
	var iterations int
	var err error

	iterations, err = strconv.Atoi(os.Args[2])
	if err != nil{
		fmt.Println("Invalid command line arguments")
		os.Exit(1)
	}

	go func() {
		defer wg.Done()
		threadName := "_Thread 1"
		result := 0
		for i := 0; i < iterations; i++ {
			if protRaceCondition {
				// Protected: use a mutex to access sumProtected
				sumProtected = calculateProtected(threadName, sumProtected, 1, 2)
			} else {
				// Unprotected: without protection, use sumUnprotected
				sumUnprotected = calculateUnprotected(threadName, sumUnprotected, 1, 2)
			}
			result = sumProtected
		}
		result1 <- result
	}()

	go func() {
		defer wg.Done()
		threadName := "Thread 2"
		result := 0
		for i := 0; i < iterations; i++ {
			if protRaceCondition {
				// Protected: use a mutex to access sumProtected
				sumProtected = calculateProtected(threadName, sumProtected, 3, 4)
			} else {
				// Unprotected: without protection, use sumUnprotected
				sumUnprotected = calculateUnprotected(threadName, sumUnprotected, 3, 4)
			}
			result = sumProtected
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

var sumMutex sync.Mutex

func calculateProtected(threadName string, previousResult, a, b int) int {
	// Lock
	sumMutex.Lock()
	defer sumMutex.Unlock()

	x := previousResult + a + b
	fmt.Printf("%s: %d\n", threadName, x)
	sleep()
	return x
}

func calculateUnprotected(threadName string, previousResult, a, b int) int {
	x := previousResult + a + b
	fmt.Printf("%s: %d\n", threadName, x)
	sleep()
	return x
}

func sleep() {
	time.Sleep(delay)
}
