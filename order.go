package main

type order struct {
	id     int
	amount float64
	order  orderStatus
}
type orderStatus int

const (
	none orderStatus = iota
	new
	received
	reserved
	filled
)
