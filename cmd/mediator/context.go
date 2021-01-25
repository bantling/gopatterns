// SPDX-License-Identifier: Apache-2.0

package main

// DataContext contains the contextual info for data flowing in either direction:
// protocol -> mediator -> operations
// protocol <- mediator <- operations
//
// If Data is non-nil, the client is sending data to store.
// If Data is nil, the client is requesting data to be returned.
type DataContext struct {
	Type   DataType
	ID     string
	Format DataFormat
	Data   []byte
}
