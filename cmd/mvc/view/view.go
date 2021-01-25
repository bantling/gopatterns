// SPDX-License-Identifier: Apache-2.0

package view

import (
	"html/template"
	"os"
)

// Template filename constants
const (
	customerTemplateFile = "cmd/mvc/view/customer.html"
)

// Template variables
var (
	customerTemplate = template.Must(template.ParseFiles(customerTemplateFile))
)

// View provides the view
type View struct{}

// NewView constructs a view
func NewView() *View {
	return &View{}
}

// RenderCustomer
func (v View) RenderCustomer(c Customer) {
	if err := customerTemplate.Execute(os.Stdout, c); err != nil {
		panic(err)
	}
}
