package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	Raw_Order_Converter()
	fmt.Println(orders)
}

func Raw_Order_Converter() {

	for _, rawOrder := range raw_Orders {
		var newOrder order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			log.Print(err)
			continue
		}
		orders = append(orders, newOrder)
	}
}

var raw_Orders = []string{
	`{"id":1, "amount":100.50, "status":1}`,
	`{"id":2, "amount":200.75, "status":2}`,
	`{"id":3, "amount":150.00, "status":3}`,
	`{"id":4, "amount":300.20, "status":4}`,
	`{"id":5, "amount":250.00, "status":0}`,
}
