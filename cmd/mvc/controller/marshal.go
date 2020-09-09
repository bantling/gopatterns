package controller

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
)

// DataFormat is the type of data being received or sent
type DataFormat uint

// String and JSON are the only supported data types
const (
	GOB DataFormat = iota
	JSON
)

var (
	dataFormatToString = map[DataFormat]string{
		GOB:  "gob",
		JSON: "json",
	}

	stringToDataFormat = map[string]DataFormat{
		"gob":  GOB,
		"json": JSON,
	}
)

// String is DataFormat Stringer
func (t DataFormat) String() string {
	return dataFormatToString[t]
}

// StringToDataFormat returns the DataFormat for a string
func StringToDataFormat(str string) DataFormat {
	return stringToDataFormat[str]
}

// Marshal data of a specified type
func Marshal(data interface{}, typ DataFormat) []byte {
	switch typ {
	case GOB:
		var buf bytes.Buffer
		gob.NewEncoder(&buf).Encode(data)
		return buf.Bytes()

	case JSON:
		buf, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		return buf

	default:
		panic(fmt.Errorf("Unrecognized data type"))
	}
}

// Unmarshal data of a specified type into the target, which must be a pointer
func Unmarshal(data []byte, typ DataFormat, target interface{}) {
	switch typ {
	case GOB:
		buf := bytes.NewReader(data)
		gob.NewDecoder(buf).Decode(target)

	default:
		if err := json.Unmarshal(data, target); err != nil {
			panic(err)
		}
	}
}
