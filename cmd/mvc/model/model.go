// SPDX-License-Identifier: Apache-2.0

package model

// Model is the public abstraction
type Model struct {
	*customerModel
}

// NewModel constructs a Model
func NewModel() *Model {
	return &Model{
		customerModel: newCustomerModel(),
	}
}
