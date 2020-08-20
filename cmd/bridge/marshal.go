package main

import (
	"encoding/json"
)

// DataType is the type of data being received or sent
type DataType uint

// String and JSON are the only supported data types
const (
	String DataType = iota
	JSON
)

// Marshal data of a specified type
func Marshal(data interface{}, typ DataType) []byte {
	switch typ {
	case String:
		return []byte(data.(string))

	default:
		buf, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		return buf
	}
}

// Unmarshal data of a specified type into the target, which must be a pointer
func Unmarshal(data []byte, typ DataType, target interface{}) {
	switch typ {
	case String:
		*(target.(*string)) = string(data)

	default:
		if err := json.Unmarshal(data, target); err != nil {
			panic(err)
		}
	}
}
