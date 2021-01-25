// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"strings"
)

type Customer struct {
	FirstName string
	LastName  string
}

type AllCustomers struct {
	Customers []Customer
}

type AllCustomersVisitorInit interface {
	Init()
}

type PreAllCustomersVisitor interface {
	VisitPreAllCustomers(allCustomers AllCustomers)
}

type PreCustomerVisitor interface {
	VisitPreCustomer(customer Customer)
}

type FirstNameVisitor interface {
	VisitFirstName(firstName string)
}

type LastNameVisitor interface {
	VisitLastName(lastName string)
}

type PostCustomerVisitor interface {
	VisitPostCustomer(customer Customer)
}

type PostAllCustomersVisitor interface {
	VisitPostAllCustomers(allCustomers AllCustomers)
}

type AllCustomersVisitor interface {
	AllCustomersVisitorInit

	PreAllCustomersVisitor

	PreCustomerVisitor
	FirstNameVisitor
	LastNameVisitor
	PostCustomerVisitor

	PostAllCustomersVisitor
}

type CustomerVisitorAdapter struct {
	initVisitor AllCustomersVisitorInit

	preAllCustomersVisitor func(AllCustomers)

	preCustomerVisitor  func(Customer)
	firstNameVisitor    func(string)
	lastNameVisitor     func(string)
	postCustomerVisitor func(Customer)

	postAllCustomersVisitor func(AllCustomers)
}

func NewCustomerVisitorAdapter(visitor ...AllCustomersVisitorInit) *CustomerVisitorAdapter {
	va := &CustomerVisitorAdapter{}

	if len(visitor) > 0 {
		va.WithVisitor(visitor[0])
	}

	return va
}

func (va *CustomerVisitorAdapter) WithVisitor(visitor AllCustomersVisitorInit) {
	if visitor == nil {
		panic(fmt.Errorf("CustomerVisitorAdapter.WithVisitor: visitor cannot be nil"))
	}

	va.initVisitor = visitor

	va.preAllCustomersVisitor = func(AllCustomers) {}
	if v, ok := visitor.(PreAllCustomersVisitor); ok {
		va.preAllCustomersVisitor = v.VisitPreAllCustomers
	}

	va.preCustomerVisitor = func(Customer) {}
	if v, ok := visitor.(PreCustomerVisitor); ok {
		va.preCustomerVisitor = v.VisitPreCustomer
	}

	va.firstNameVisitor = func(string) {}
	if v, ok := visitor.(FirstNameVisitor); ok {
		va.firstNameVisitor = v.VisitFirstName
	}

	va.lastNameVisitor = func(string) {}
	if v, ok := visitor.(LastNameVisitor); ok {
		va.lastNameVisitor = v.VisitLastName
	}

	va.postCustomerVisitor = func(Customer) {}
	if v, ok := visitor.(PostCustomerVisitor); ok {
		va.postCustomerVisitor = v.VisitPostCustomer
	}

	va.postAllCustomersVisitor = func(AllCustomers) {}
	if v, ok := visitor.(PostAllCustomersVisitor); ok {
		va.postAllCustomersVisitor = v.VisitPostAllCustomers
	}
}

func (va CustomerVisitorAdapter) Init() {
	va.initVisitor.Init()
}

func (va CustomerVisitorAdapter) VisitPreAllCustomers(allCustomers AllCustomers) {
	va.preAllCustomersVisitor(allCustomers)
}

func (va CustomerVisitorAdapter) VisitPreCustomer(customer Customer) {
	va.preCustomerVisitor(customer)
}

func (va CustomerVisitorAdapter) VisitFirstName(firstName string) {
	va.firstNameVisitor(firstName)
}

func (va CustomerVisitorAdapter) VisitLastName(lastName string) {
	va.lastNameVisitor(lastName)
}

func (va CustomerVisitorAdapter) VisitPostCustomer(customer Customer) {
	va.postCustomerVisitor(customer)
}

func (va CustomerVisitorAdapter) VisitPostAllCustomers(allCustomers AllCustomers) {
	va.postAllCustomersVisitor(allCustomers)
}

type CustomersPrinter struct {
	bldr strings.Builder
}

func (p *CustomersPrinter) Init() { p.bldr.Reset() }

func (p *CustomersPrinter) VisitPreCustomer(Customer) { p.bldr.WriteRune('[') }

func (p *CustomersPrinter) VisitFirstName(firstName string) { p.bldr.WriteString(firstName) }

func (p *CustomersPrinter) VisitLastName(lastName string) {
	p.bldr.WriteString(", ")
	p.bldr.WriteString(lastName)
}

func (p *CustomersPrinter) VisitPostCustomer(Customer) { p.bldr.WriteRune(']') }

func (p *CustomersPrinter) Result() string { return p.bldr.String() }

type AllCustomersDepthFirstWalker struct {
	visitor AllCustomersVisitor
}

func NewAllCustomersDepthFirstWalker(visitor AllCustomersVisitor) AllCustomersDepthFirstWalker {
	return AllCustomersDepthFirstWalker{visitor}
}

func (w AllCustomersDepthFirstWalker) walk(allCustomers AllCustomers) {
	w.visitor.Init()

	w.visitor.VisitPreAllCustomers(allCustomers)

	for _, c := range allCustomers.Customers {
		w.visitor.VisitPreCustomer(c)
		w.visitor.VisitFirstName(c.FirstName)
		w.visitor.VisitLastName(c.LastName)
		w.visitor.VisitPostCustomer(c)
	}

	w.visitor.VisitPostAllCustomers(allCustomers)
}

func main() {
	allCustomers := AllCustomers{
		Customers: []Customer{
			Customer{FirstName: "John", LastName: "Doe"},
			Customer{FirstName: "Jane", LastName: "Doe"},
		},
	}

	p := &CustomersPrinter{}
	va := NewCustomerVisitorAdapter(p)
	w := NewAllCustomersDepthFirstWalker(va)
	w.walk(allCustomers)
	fmt.Println(p.Result())
}
