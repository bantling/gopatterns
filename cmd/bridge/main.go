// SPDX-License-Identifier: Apache-2.0

package main

import (
	"time"
)

func main() {
	traffic := NewTraffic()
	ftp := NewFTPTraffic(traffic)
	http := NewHTTPTraffic(traffic)

	cust := Customer{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Address: Address{
			Line:     "123 Sesame St",
			City:     "New York",
			Region:   "New York",
			MailCode: "12345",
		},
	}

	buf := Marshal(cust, GOB)
	ftp.Receive("customer.gob", buf)
	http.Send("/customer/1", JSON)

	invoice := Invoice{
		Number:     "A14",
		CustomerID: 1,
		Date:       time.Now(),
		Lines: []Line{
			Line{
				Product:  "Apples",
				Price:    "3.10",
				Qty:      5,
				Extended: "15.50",
			},
		},
	}

	buf = Marshal(invoice, JSON)
	http.Receive("/invoice/A14", JSON, buf)

	ftp.Send("1.json", InvoiceType)
}
