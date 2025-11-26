package main

import "fmt"

type order struct {
	ID     int
	Amount float64
	Status orderStatus
}

func (o order) String() string {
	return fmt.Sprintf("Order ID:%v, Amount:%v, Status :%v", o.ID, o.Amount,
		orderStatusToText(o.Status))
}

type orderStatus int

const (
	none orderStatus = iota
	new
	received
	reserved
	filled
)

func orderStatusToText(status orderStatus) string {
	switch status {
	case none:
		return "None"
	case new:
		return "New"
	case received:
		return "Received"
	case reserved:
		return "Reserved"
	case filled:
		return "Filled"
	default:
		return "Unknown Status"
	}

}

var orders = []order{}
