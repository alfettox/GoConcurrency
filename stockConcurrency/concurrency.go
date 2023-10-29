package main

import (
	"fmt"
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup
	var mu sync.Mutex

	waitGroup.Add(2)

	appleStock := 0   // Number of Apple stock options
	amazonStock := 0  // Number of Amazon stock options

	go func() {
		defer waitGroup.Done() //execute later
		trackStockOptions("Martin", "Apple", &appleStock, &mu)
	}()

	go func() {
		defer waitGroup.Done()
		trackStockOptions("Louise", "Amazon", &amazonStock, &mu)
	}()

	stockNoteChannel := make(chan string)

	go func() {
		leaveNote("Martin", "I bought Apple stock options.", stockNoteChannel)
	}()
	leaveNote("Louise", "I bought Amazon stock options.", stockNoteChannel)

	waitGroup.Wait()
	fmt.Printf("Both workers have stock options:\nApple: %d options\nAmazon: %d options\n", appleStock, amazonStock)
}

func trackStockOptions(workerName string, companyName string, stockOptions *int, mu *sync.Mutex) {
	for i := 1; i <= 10; i++ {
		mu.Lock()
		*stockOptions++
		mu.Unlock()
		fmt.Printf("%s: %s stock options: %d options\n", workerName, companyName, *stockOptions)
	}
}

func leaveNote(workerName string, note string, channel chan string) {
	note = workerName + ": " + note
	channel <- note
	fmt.Println(note)
}
