// SPDX-License-Identifier: Apache-2.0

package main

// Create all instances
var (
	ftpTraffic         = NewFTPTraffic()
	httpTraffic        = NewHTTPTraffic()
	mediator           = NewMediator()
	customerOperations = NewCustomerOperations()
	invoiceOperations  = NewInvoiceOperations()
)

// Wire them up to refer to each other
func init() {
	ftpTraffic.mediator = mediator
	httpTraffic.mediator = mediator
	mediator.ftpTraffic = ftpTraffic
	mediator.httpTraffic = httpTraffic

	mediator.customerOps = customerOperations
	mediator.invoiceOps = invoiceOperations
}
