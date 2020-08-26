package main

import (
	"time"
)

type Customer struct {
	ID        int
	FirstName string
	LastName  string
	Address   Address
}

type Address struct {
	Line     string
	City     string
	Region   string
	Country  string
	MailCode string
}

type Invoice struct {
	Number     string
	CustomerID int
	Date       time.Time
	Lines      []Line
}

type Line struct {
	Product  string
	Price    string
	Qty      uint
	Extended string
}
