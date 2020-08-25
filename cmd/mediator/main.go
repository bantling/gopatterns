package main

import (
	"time"
)

func main() {
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
	ftpTraffic.Request("/customer/1.gob", buf)

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
	httpTraffic.Request("/invoice/A14", JSON, buf)

	ftpTraffic.Request("/customer/1.json", nil)
	httpTraffic.Request("/invoice/1", GOB, nil)
}
