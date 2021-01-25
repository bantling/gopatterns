// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"fmt"
	"io"
)

// LineCounter is an io.Reader that counts lines of another reader
type LineCounter struct {
	lineNo int
	reader io.Reader
	lastCR bool
}

// NewLineCounter constructs a LineCounter
func NewLineCounter(r io.Reader) *LineCounter {
	return &LineCounter{lineNo: 1, reader: r, lastCR: false}
}

// Read is the io.Reader interface
func (lc *LineCounter) Read(b []byte) (n int, err error) {
	// Call underlying reader to get results to return into named return parameters
	n, err = lc.reader.Read(b)

	// Scan the slice from [0,n)
	for i := 0; i < n; i++ {
		// Count CR, LF,or CRLF as one newline
		switch b[i] {
		case '\r':
			lc.lineNo++
			// Mark last char was a CR in case next char is an LF.
			// This works even if CR is read as last byte in one Read, and LF is first byte in next Read.
			// A sequence of CRs will increment on each CR, that is correct, a lone CR is the old MacOS line ending.
			lc.lastCR = true
		case '\n':
			if lc.lastCR {
				// CRLF sequence, already incremented
				lc.lastCR = false
			} else {
				// LF not preceded by CR, increment
				lc.lineNo++
			}
		default:
			// Turn off lastCR so it doesn't wait until the next LF and incorrectly assume it's a CRLF
			lc.lastCR = false
		}
	}

	return
}

func main() {
	// Create a slice of bytes to treat as our input
	data := []byte("Line 1\rLine2\nLine3\r\nLine4")

	// Try reading one char at a time, then 4 chars, then 100 chars (more than line length)
	for _, bufSize := range []int{1, 4, 100} {
		// Make a new Reader of the data
		rdr := NewLineCounter(bytes.NewReader(data))

		// Make a buffer of the current size
		buf := make([]byte, bufSize)

		// Read until EOF
		for _, err := rdr.Read(buf); err != io.EOF; _, err = rdr.Read(buf) {
			if err != nil {
				panic(err)
			}
		}

		// Dump line counter
		fmt.Printf("Number of lines = %d\n", rdr.lineNo)
	}
}
