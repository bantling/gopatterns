package main

import (
	"fmt"
	"strconv"
)

// Mediator that handles communication between Receiver/Sender and CustomerOperations/InvoiceOperations
type Mediator struct {
	ftpTraffic  *FTPTraffic
	httpTraffic *HTTPTraffic
	customerOps *CustomerOperations
	invoiceOps  *InvoiceOperations
}

// NewMediator constructs a Mediator
func NewMediator() *Mediator {
	return &Mediator{}
}

// Perform the operation
func (t *Mediator) Perform(ctx DataContext) DataContext {
	responseCtx := ctx

	switch ctx.Type {
	case CustomerType:
		if ctx.Data != nil {
			// Store
			var cust Customer
			Unmarshal(ctx.Data, ctx.Format, &cust)
			t.customerOps.SetCustomer(cust)
			responseCtx.Data = nil
		} else {
			// Retrieve
			id, _ := strconv.Atoi(ctx.ID)
			cust := t.customerOps.GetCustomer(id)
			responseCtx.Data = Marshal(cust, ctx.Format)
		}

	case InvoiceType:
		if ctx.Data != nil {
			// Store
			var invoice Invoice
			Unmarshal(ctx.Data, ctx.Format, &invoice)
			t.invoiceOps.SetInvoice(invoice)
			responseCtx.Data = nil
		} else {
			// Retrieve
			id, _ := strconv.Atoi(ctx.ID)
			invoices := t.invoiceOps.GetInvoicesForCustomer(id)
			responseCtx.Data = Marshal(invoices, ctx.Format)
		}

	default:
		panic(fmt.Errorf("Unknown request type"))
	}

	return responseCtx
}
