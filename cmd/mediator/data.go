package main

import (
	"time"
)

type DataType uint

const (
	CustomerType DataType = iota
	InvoiceType
)

var (
	dataTypeToString = map[DataType]string{
		CustomerType: "customer",
		InvoiceType:  "invoice",
	}

	stringToDataType = map[string]DataType{
		"customer": CustomerType,
		"invoice":  InvoiceType,
	}
)

// String is DataType Stringer
func (dt DataType) String() string {
	return dataTypeToString[dt]
}

// StringToDataType returns the DataType for a string
func StringToDataType(str string) DataType {
	return stringToDataType[str]
}

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
