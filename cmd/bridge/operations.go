package main

// CustomerOperations contains operations on Customer
type CustomerOperations struct {
	Customers map[int]Customer
}

// SetCustomer adds or replaces a customer by id
func (c *CustomerOperations) SetCustomer(data Customer) {
	c.Customers[data.ID] = data
}

// GetCustomer returns one customer by id and a flag indicating whether or not the customer exists
func (c CustomerOperations) GetCustomer(id int) (Customer, bool) {
	cust, exists := c.Customers[id]
	return cust, exists
}

// InvoiceOperations contains operations on Invoices
type InvoiceOperations struct {
	Invoices       map[string]Invoice
	InvoicesByCust map[int][]Invoice
}

// SetInvoice adds or replaces an invoice by number
func (i InvoiceOperations) SetInvoice(data Invoice) {
	i.Invoices[data.Number] = data
	i.InvoicesByCust[data.Customer.ID] = append(i.InvoicesByCust[data.Customer.ID], data)
}

// GetInvoicesForCustomer gets all invoices for a given Customer id
func (i InvoiceOperations) GetInvoicesForCustomer(id int) []Invoice {
	return i.InvoicesByCust[id]
}
