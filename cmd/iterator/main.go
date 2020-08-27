package main

import (
	"fmt"
	"reflect"
)

// ProductNames contains a set of product names
type ProductNames struct {
	names map[string]bool
	modifications uint
}

// NewProductNames constructs a ProductNames
func NewProductNames(products ...string) *ProductNames {
	m := map[string]bool{}
	for _, product := range products {
		m[product] = true
	}
	
	return &ProductNames{
		names: m,
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

// Iter generates an iterator of the product names, and panics if modifications were made since the iterator was created
// Although it makes no modifications it has to have a pointer receiver to see updates to the modification counter
func (p *ProductNames) Iter() func() (product string, valid bool) {
	var (
		modifications = p.modifications
		mapIter = reflect.ValueOf(p.names).MapRange()
	)
	return func() (string, bool) {
		// Die if modifications made since last call
		if modifications != p.modifications {
			panic(fmt.Errorf("The set of product names has been modified since the iterator was created"))
		}
		
		// Return next product if there is one
		if mapIter.Next() {
			return mapIter.Key().String(), true
		}
		
		// Return empty string, false if there isn't
		return "", false
	}
}

func main() {
	p := NewProductNames("car", "bicycle", "tvs")
	iter := p.Iter()
	for product, hasNext := iter(); hasNext; product, hasNext = iter() {
		fmt.Println("Product ", product)
	}
	
	// Die if modifications made during iteration
	iter = p.Iter()
	p.RemoveProductName("tvs")
	iter()
}
