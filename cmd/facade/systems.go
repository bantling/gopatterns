package main

import (
	"time"
)

// CustomerService accesses customer records
type CustomerService struct {
	customers map[int]Customer
}

// NewCustomerService constructs a CustomerService
func NewCustomerService() *CustomerService {
	return &CustomerService{
		customers: map[int]Customer{
			1: Customer{
				ID:        1,
				FirstName: "John",
				LastName:  "Doe",
				Address: Address{
					Line:     "123 Sesame St",
					City:     "New York",
					Region:   "New York",
					Country:  "USA",
					MailCode: "12345",
				},
			},
		},
	}
}

// GetCustomer returns a customer by id
func (cs *CustomerService) GetCustomer(id int) Customer {
	return cs.customers[id]
}

// InvoiceService accesses invoice records
type InvoiceService struct {
	invoicesByCustomer map[int][]Invoice
}

// NewInvoiceService constructs an InvoiceService
func NewInvoiceService() *InvoiceService {
	return &InvoiceService{
		invoicesByCustomer: map[int][]Invoice{
			1: []Invoice{
				Invoice{
					Number:     "A14",
					CustomerID: 1,
					Date:       time.Now(),
					Lines: []Line{
						Line{
							Product:  "Apples",
							Price:    "2.50/lb",
							Qty:      3,
							Extended: "7.50",
						},
					},
				},
			},
		},
	}
}

// GetInvoicesForCustomer retrieves all invoices for a given customer id
func (is *InvoiceService) GetInvoicesForCustomer(id int) []Invoice {
	return is.invoicesByCustomer[id]
}
