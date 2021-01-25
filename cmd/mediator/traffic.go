// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"strings"
)

// FTPTraffic represents data to be received/returned via FTP
type FTPTraffic struct {
	mediator *Mediator
}

// NewFTPTraffic constructs FTPTraffic
func NewFTPTraffic() *FTPTraffic {
	return &FTPTraffic{}
}

// Request is called by a virtual FTP server when it receives a request to store or retrieve data
func (f FTPTraffic) Request(path string, data []byte) {
	typeIDAndFormat := strings.Split(path, ".")
	typeAndID := strings.Split(typeIDAndFormat[0], "/")

	// Leading / so index 0 is empty string
	typ := StringToDataType(typeAndID[1])
	id := typeAndID[2]
	format := StringToDataFormat(typeIDAndFormat[1])

	fmt.Printf("Requesting FTP for %s %s: %s = %s\n", typ, id, format, data)
	responseCtx := f.mediator.Perform(DataContext{Type: typ, ID: id, Format: format, Data: data})
	if responseCtx.Data == nil {
		// Successful upload
		fmt.Println("FTP Successful Upload to", typ, id)
	} else {
		// Successful download
		fmt.Printf("FTP Successful Download of %s %s: %s = %s\n", typ, id, responseCtx.Format, responseCtx.Data)
	}
}

// HTTPTraffic represents data to be sent/received via HTTP
type HTTPTraffic struct {
	mediator *Mediator
}

// NewHTTPTraffic constructs HTTPTraffic
func NewHTTPTraffic() *HTTPTraffic {
	return &HTTPTraffic{}
}

// Request is called by an HTTP server when it receives a request to store or retrieve data
func (h HTTPTraffic) Request(path string, format DataFormat, data []byte) {
	typeAndID := strings.Split(path, "/")

	// Leading / so index 0 is empty string
	typ := StringToDataType(typeAndID[1])
	id := typeAndID[2]

	fmt.Printf("Requesting HTTP for %s %s: %s = %s\n", typ, id, format, data)
	responseCtx := h.mediator.Perform(DataContext{Type: typ, ID: id, Format: format, Data: data})
	if responseCtx.Data == nil {
		// Successful upload
		fmt.Println("HTTP Successful Upload to", typ, id)
	} else {
		// Successful download
		fmt.Printf("HTTP Successful Download of %s %s: %s = %s\n", typ, id, responseCtx.Format, responseCtx.Data)
	}
}
