// SPDX-License-Identifier: Apache-2.0

package model

type Customer struct {
	ID        int
	FirstName string
	LastName  string
	Address   Address
}

type Address struct {
	Line     string
	City     string
	Region   string
	Country  string
	MailCode string
}
