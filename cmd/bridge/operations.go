// SPDX-License-Identifier: Apache-2.0

package main

// CustomerOperations contains operations on Customer
type CustomerOperations struct {
	customers map[int]Customer
}

// NewCustomerOperations constructs CustomerOperations
func NewCustomerOperations() *CustomerOperations {
	return &CustomerOperations{customers: map[int]Customer{}}
}

// SetCustomer adds or replaces a customer by id
func (c *CustomerOperations) SetCustomer(data Customer) {
	c.customers[data.ID] = data
}

// GetCustomer returns one customer by id and a flag indicating whether or not the customer exists
func (c CustomerOperations) GetCustomer(id int) Customer {
	return c.customers[id]
}

// InvoiceOperations contains operations on Invoices
type InvoiceOperations struct {
	invoices         map[string]Invoice
	invoicesByCustID map[int][]Invoice
}

// NewInvoiceOperations constructs InvoiceOperations
func NewInvoiceOperations() *InvoiceOperations {
	return &InvoiceOperations{
		invoices:         map[string]Invoice{},
		invoicesByCustID: map[int][]Invoice{},
	}
}

// SetInvoice adds or replaces an invoice by number
func (i *InvoiceOperations) SetInvoice(data Invoice) {
	i.invoices[data.Number] = data
	i.invoicesByCustID[data.CustomerID] = append(i.invoicesByCustID[data.CustomerID], data)
}

// GetInvoicesForCustomer gets all invoices for a given Customer id
func (i InvoiceOperations) GetInvoicesForCustomer(id int) []Invoice {
	return i.invoicesByCustID[id]
}
