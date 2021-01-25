// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"fmt"

	"github.com/bantling/gopatterns/cmd/mvc/model"
	"github.com/bantling/gopatterns/cmd/mvc/view"
)

// controller mediates between the model and view
type controller struct {
	model *model.Model
	view  *view.View
}

// NewController constructs a Controller
func NewController() *controller {
	return &controller{
		model: model.NewModel(),
		view:  view.NewView(),
	}
}

// GetCustomer retrieves a customer and renders the result
func (c controller) GetCustomer(id int) {
	fmt.Println("==== SetCustomer(Customer)")
	c.model.SetCustomer(
		model.Customer{
			ID:        1,
			FirstName: "John",
			LastName:  "Doe",
			Address: model.Address{
				Line:     "123 Sesame St",
				City:     "New York",
				Region:   "New York",
				Country:  "USA",
				MailCode: "12345",
			},
		},
	)

	fmt.Printf("==== GetCustomer(%d)\n", id)
	mcust := c.model.GetCustomer(id)
	vcust := view.Customer{
		ID:        mcust.ID,
		FirstName: mcust.FirstName,
		LastName:  mcust.LastName,
		Address:   view.Address(mcust.Address),
	}
	c.view.RenderCustomer(vcust)
}
