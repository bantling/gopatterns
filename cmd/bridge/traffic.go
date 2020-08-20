package main

import (
	"strings"
)

// Receiver receives data
type Receiver interface {
	Receive() (path string, typ DataType, data []byte)
}

// Sender sends data
type Sender interface {
	Send(path string, typ DataType, data []byte)
}

// FTPTraffic represents data to be received via FTP
type FTPTraffic struct {
	Filename string
	Type     DataType
	Data     []byte
}

// Receive is called by a virtual FTP server when it receives an upload
func (f *FTPTraffic) Receive(filename string, typ DataType, data []byte) {
	f.Filename, f.Type, f.Data = filename, typ, data
}

// HTTPTraffic represents data to be sent/received via HTTP
type HTTPTraffic struct {
	Filename string
	Type     DataType
	Data     []byte
}

// Receive is called when by an HTTP server when it receives data in the request body
func (h HTTPTraffic) Receive(path string, typ DataType, data []byte) {
	pathParts := strings.Split(path, "/")
	h.Filename, h.Type, h.Data = pathParts[len(pathParts)-1], typ, data
}

// Send returns data to an HTTP client in the response body
func (h HTTPTraffic) Send(path string, typ DataType, data []byte) {
	h.Receive(path, typ, data)
}
