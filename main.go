package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

func main() {
	// Using WaitGroup to wait for goroutine to finish
	// Note :- suitable for multiple goroutines
	// if you will call it multiple times race condition may occur on orders slice
	// so proper synchronization is needed in that case and for that mutex can be used

	/*
		//here the variable is defined at top level and the method is tightly coupled with it.
		var wg sync.WaitGroup
		wg.Add(3)
		go func() {
			defer wg.Done()
			log.Println("Starting Raw Order Conversion in Goroutine")
			Raw_Order_Converter()
		}()
		wg.Wait()
	*/

	/* Using channel to signal completion of goroutine
	 Note :- only suitable for single goroutine
	done := make(chan bool)
	go func() {
		Raw_Order_Converter()
		done <- true
	}()
	log.Println("Waiting for Goroutine to finish")
	// this can also be called directly without goroutine

	//Raw_Order_Converter()

	*/

	//to avoid global variable coupling we can pass the waitgroup as parameter to the function
	var wg sync.WaitGroup
	wg.Add(1)
	Raw_Order_Converter(&wg)
	wg.Wait()
	log.Println("Goroutine finished execution")
	fmt.Println(orders)
}

func Raw_Order_Converter(wg *sync.WaitGroup) {

	for _, rawOrder := range raw_Orders {
		var newOrder order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			log.Print(err)
			continue
		}
		orders = append(orders, newOrder)
	}
	wg.Done()

}

var raw_Orders = []string{
	`{"id":1, "amount":100.50, "status":1}`,
	`{"id":2, "amount":200.75, "status":2}`,
	`{"id":3, "amount":150.00, "status":3}`,
	`{"id":4, "amount":300.20, "status":4}`,
	`{"id":5, "amount":250.00, "status":0}`,
}
