// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
)

func main() {
	facadeSvc := NewFacadeService()
	fmt.Printf("customerInvoices for id 1 = %+v\n", facadeSvc.GetCustomerInvoicesByID(1))
}
