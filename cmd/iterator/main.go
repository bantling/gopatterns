// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"reflect"
)

// ProductNames contains a set of product names
type ProductNames struct {
	names         map[string]bool
	modifications uint
}

// ProductNamesIterator is an iterator of ProductNames
// Call Next() to advance to next value, and only if it returns true, call Name() to get next name
type ProductNamesIterator struct {
	p             *ProductNames
	modifications uint
	iter          *reflect.MapIter
}

// Next advances to next product name, returning true if there are any more left to iterate
// Panics if changes have been made since iterator was created
func (pni *ProductNamesIterator) Next() bool {
	if pni.modifications != pni.p.modifications {
		panic(fmt.Errorf("The set of product names has been modified since the iterator was created"))
	}

	return pni.iter.Next()
}

// Name returns name of next product
// Panics if changes have been made since iterator was created, or iterator has been exhausted
func (pni *ProductNamesIterator) Name() string {
	if pni.modifications != pni.p.modifications {
		panic(fmt.Errorf("The set of product names has been modified since the iterator was created"))
	}

	return pni.iter.Key().String()
}

// NewProductNames constructs a ProductNames
func NewProductNames(products ...string) *ProductNames {
	m := map[string]bool{}
	for _, product := range products {
		m[product] = true
	}

	return &ProductNames{
		names:         m,
		modifications: 0,
	}
}

// AddProductName builder adds a product name
func (p *ProductNames) AddProductName(product string) *ProductNames {
	p.names[product] = true
	p.modifications++
	return p
}

// RemoveProductName builder removes a product name
func (p *ProductNames) RemoveProductName(product string) *ProductNames {
	delete(p.names, product)
	p.modifications++
	return p
}

// Iter creates an iterator of the product names
func (p *ProductNames) Iter() *ProductNamesIterator {
	return &ProductNamesIterator{
		p:             p,
		modifications: p.modifications,
		iter:          reflect.ValueOf(p.names).MapRange(),
	}
}

func main() {
	p := NewProductNames("car", "bicycle", "tvs")
	for productIter := p.Iter(); productIter.Next(); {
		fmt.Println("Product ", productIter.Name())
	}

	// Die if modifications made during iteration
	productIter := p.Iter()
	p.RemoveProductName("tvs")
	productIter.Next()
}
