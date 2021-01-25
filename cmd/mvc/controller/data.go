// SPDX-License-Identifier: Apache-2.0

package controller

// DataType describes all types of data
type DataType uint

// DataType constants
const (
	CustomerType DataType = iota
)

var (
	dataTypeToString = map[DataType]string{
		CustomerType: "customer",
	}

	stringToDataType = map[string]DataType{
		"customer": CustomerType,
	}
)

// String is DataType Stringer
func (dt DataType) String() string {
	return dataTypeToString[dt]
}

// StringToDataType returns the DataType for a string
func StringToDataType(str string) DataType {
	return stringToDataType[str]
}
