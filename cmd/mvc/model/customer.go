package model

// customerModel contains operations on Customer
type customerModel struct {
	customers map[int]Customer
}

// newCustomerModel constructs customerModel
func newCustomerModel() *customerModel {
	return &customerModel{customers: map[int]Customer{}}
}

// SetCustomer adds or replaces a customer by id
func (c *customerModel) SetCustomer(data Customer) {
	c.customers[data.ID] = data
}

// GetCustomer returns one customer by id and a flag indicating whether or not the customer exists
func (c customerModel) GetCustomer(id int) Customer {
	return c.customers[id]
}
