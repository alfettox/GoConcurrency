package main

import (
	"fmt"
	"sync"
	"time"
	"strings"
)

func main() {
	var waitGroup sync.WaitGroup
	var mu sync.Mutex

	waitGroup.Add(2)

	ericssonStock := 0 // Total Ericsson stocks
	huaweiStock := 0   // Total Huawei stocks

	stockNoteChannel := make(chan string)

	// NBC's goroutine
	go func() {
		defer waitGroup.Done()
		trackStockOptions("NBC", "Huawei", &huaweiStock, &mu, stockNoteChannel)
	}()

	// Morgan Stanley's goroutine
	go func() {
		defer waitGroup.Done()
		trackStockOptions("Morgan Stanley", "Ericsson", &ericssonStock, &mu, stockNoteChannel)
	}()

	go func() {
		waitGroup.Wait()
		close(stockNoteChannel)
	}()

	fmt.Printf("%-20s %-20s %-10s\n", "Bank", "Company", "Stock Options")
	fmt.Println(strings.Repeat("-", 50))

	for message := range stockNoteChannel {
		fmt.Println(message)
	}

	fmt.Println(strings.Repeat("-", 50))
	fmt.Printf("%-20s %-20s %-10s\n", "Total", "Ericsson", "Huawei")
	fmt.Println("_________________________________________________")
	fmt.Printf("%-20s %-20d %-10d\n", "TOTAL STOCKS: ", ericssonStock, huaweiStock)
}

func trackStockOptions(bankName string, companyName string, stockOptions *int, mu *sync.Mutex, channel chan string) {
	for i := 1; i <= 10; i++ {
		mu.Lock()
		*stockOptions++
		mu.Unlock()
		message := fmt.Sprintf("%-20s %-20s %-10d", bankName, companyName, *stockOptions)
		time.Sleep(100 * time.Millisecond)
		channel <- message
	}
}
