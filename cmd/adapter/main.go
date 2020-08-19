package main

import (
	"bytes"
	"fmt"
	"io"
)

const (
	// bufferSize is the size of the buffer that is guaranteed to be large enough to ready any single UTF-8 character
	bufferSize = 16
)

// RuneReaderAdapter adapts a Reader into a RuneReader
type RuneReaderAdapter struct {
	br      io.Reader
	buf     []byte
	bufSize int
	rr      *bytes.Reader
}

// NewRuneReaderAdapter constructs a RuneReaderAdapter
func NewRuneReaderAdapter(r io.Reader) *RuneReaderAdapter {
	buffer := make([]byte, bufferSize)
	return &RuneReaderAdapter{
		br:      r,
		buf:     buffer,
		bufSize: 0,
		rr:      bytes.NewReader(buffer),
	}
}

// ReadRune is the RuneReader interface
func (a *RuneReaderAdapter) ReadRune() (ch rune, runeSize int, err error) {
	// If bufSize = -1, nothing left to read
	if a.bufSize == -1 {
		return 0, 0, io.EOF
	}

	// Read buffer size - bufSize bytes at end of buf from underlying Reader
	// This fills in the remaining bytes left over from last read
	// The first read will fill the whole buffer
	subBuf := a.buf[a.bufSize:]
	bytesRead, err := a.br.Read(subBuf)

	// Stop on problem errors
	if (err != nil) && (err != io.EOF) {
		return 0, 0, err
	}

	// Increase bufSize by the number of bytes read, which could be zero
	a.bufSize += bytesRead

	// Reset the underlying RuneReader to the beginning of the rune buffer
	a.rr.Seek(0, io.SeekStart)

	// Read next rune from beginning of buffer
	// This sets the return values, so only need naked returns going forward
	ch, runeSize, err = a.rr.ReadRune()

	// Stop on problem errors
	if (err != nil) && (err != io.EOF) {
		return
	}

	// Copy the bytes after the rune backwards in the buffer
	copy(a.buf[:bufferSize-runeSize], a.buf[runeSize:])

	// Reduce bufSize by runeSize
	a.bufSize -= runeSize

	// If the new bufSize is 0, there are no more runes to read
	if a.bufSize == 0 {
		a.bufSize = -1
	}

	// Return the rune, size, and error we just read
	return
}

func dump(input string) {
	rra := NewRuneReaderAdapter(bytes.NewReader([]byte(input)))

	for {
		ch, size, err := rra.ReadRune()
		fmt.Printf("%s, %d, %v\n", string(ch), size, err)
		if err == io.EOF {
			break
		}
	}
}

func main() {
	// Single byte chars
	dump("A")
	dump("AB")
	//    123456789012345
	dump("ABCDEFGHIJKLMNO")
	//   1234567890123456
	dump("ABCDEFGHIJKLMNOP")
	//   1234567890123457
	dump("ABCDEFGHIJKLMNOPQ")

	// Multibyte chars
	dump("£¥§©Æ的的Æ©§¥£")
}
