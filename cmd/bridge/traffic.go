package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Traffic abstract struct that handles common functionality of Receiver and Sender
type Traffic struct {
	*CustomerOperations
	*InvoiceOperations
}

// NewTraffic constructs Traffic
func NewTraffic() *Traffic {
	return &Traffic{
		CustomerOperations: NewCustomerOperations(),
		InvoiceOperations:  NewInvoiceOperations(),
	}
}

// Receive is called by implementer
func (t *Traffic) Receive(filename string, data []byte) {
	parts := strings.Split(filename, ".")
	name := parts[0]
	ext := parts[1]
	typ := StringToDataFormat(ext)
	fmt.Printf("Received file of type %s in %s format\n", name, typ) 
	
	switch name {
	case "customer":
		var cust Customer
		Unmarshal(data, typ, &cust)
		t.SetCustomer(cust)

	case "invoice":
		var invoice Invoice
		Unmarshal(data, typ, &invoice)
		t.SetInvoice(invoice)

	default:
		panic(fmt.Errorf("Unknown receive path"))
	}
}

// Send is called by implementer
func (t Traffic) Send(filename string, dt DataType) {
	parts := strings.Split(filename, ".")
	name := dt.String()
	ext := parts[1]
	typ := StringToDataFormat(ext)

	switch name {
	case "customer":
		id, _ := strconv.Atoi(filename)
		cust := t.GetCustomer(id)
		send := Marshal(cust, typ)
		fmt.Printf("Sending Customer %s\n", send)

	case "invoice":
		id, _ := strconv.Atoi(parts[0])
		invoices := t.GetInvoicesForCustomer(id)
		send := Marshal(invoices, typ)
		fmt.Printf("Sending Invoices for Invoice %d: %s\n", id, send)

	default:
		panic(fmt.Errorf("Unknown send path"))
	}
	
	fmt.Printf("Sent file %s of type %s in %s format\n", filename, name, typ)
}

// FTPTraffic represents data to be received via FTP
type FTPTraffic struct {
	traffic *Traffic
}

// NewFTPTraffic constructs FTPTraffic
func NewFTPTraffic(t *Traffic) *FTPTraffic {
	return &FTPTraffic{traffic: t}
}

// Receive is called by a virtual FTP server when it receives an upload
func (f *FTPTraffic) Receive(filename string, data []byte) {
	fmt.Printf("Receiving FTP for %s: %s\n", filename, data)
	f.traffic.Receive(filename, data)
}

// Send is called by a virtual FTP server when it receives a download
func (f *FTPTraffic) Send(filename string, dt DataType) {
	fmt.Println("Sending FTP for", filename)
	f.traffic.Send(filename, dt)
}

// HTTPTraffic represents data to be sent/received via HTTP
type HTTPTraffic struct {
	traffic *Traffic
}

// NewHTTPTraffic constructs HTTPTraffic
func NewHTTPTraffic(t *Traffic) *HTTPTraffic {
	return &HTTPTraffic{traffic: t}
}

// Receive is called when by an HTTP server when it receives data in the request body
func (h HTTPTraffic) Receive(path string, typ DataFormat, data []byte) {
	pathParts := strings.Split(path, "/")
	filename := pathParts[1]
	fmt.Printf("Receiving HTTP for %s: %s\n", filename, data)
	h.traffic.Receive(filename+"."+typ.String(), data)
}

// Send returns data to an HTTP client in the response body
func (h HTTPTraffic) Send(path string, typ DataFormat) {
	pathParts := strings.Split(path, "/")
	filename := pathParts[1]
	fmt.Println("Sending HTTP for", filename)
	h.traffic.Send(filename + "." + typ.String(), StringToDataType(filename))
}
