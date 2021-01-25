// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/bantling/gopatterns/cmd/mvc/controller"
)

func main() {
	cntl := controller.NewController()
	cntl.GetCustomer(1)
}
