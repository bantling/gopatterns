package main

import (
	"time"
)

// FacadeInvoice is a simplkified invoice - it doesn't need a customer id
type FacadeInvoice struct {
	Number string
	Date   time.Time
	Lines  []Line
}

// FacadeCustomerInvoices summarizes the result of multiple operations
type FacadeCustomerInvoices struct {
	Customer Customer
	Invoices []FacadeInvoice
}

// FacadeService is the facade service
type FacadeService struct {
	customerSvc *CustomerService
	invoiceSvc  *InvoiceService
}

// NewFacadeService constructs a FacadeService
func NewFacadeService() *FacadeService {
	return &FacadeService{
		customerSvc: NewCustomerService(),
		invoiceSvc:  NewInvoiceService(),
	}
}

// GetCustomerInvoicesByID queries the customer and invoice services
func (f FacadeService) GetCustomerInvoicesByID(id int) FacadeCustomerInvoices {
	customer := f.customerSvc.GetCustomer(id)
	invoices := f.invoiceSvc.GetInvoicesForCustomer(id)

	facadeInvoices := make([]FacadeInvoice, len(invoices))
	for i, invoice := range invoices {
		facadeInvoices[i] = FacadeInvoice{
			Number: invoice.Number,
			Date:   invoice.Date,
			Lines:  invoice.Lines,
		}
	}

	return FacadeCustomerInvoices{
		Customer: customer,
		Invoices: facadeInvoices,
	}
}
