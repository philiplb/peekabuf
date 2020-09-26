/*
 * This file is part of the Peek-A-Buf package.
 *
 * (c) Philip Lehmann-Böhm <philip@philiplb.de>
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

// Package peekabuf offers buffers with the capability to be peeked without
// side effects on the unread functionality. Along with that, it returns a
// special `EOF` rune instead of an error if it reaches the end of the buffer.
package peekabuf

import (
	"bufio"
	"container/list"
	"io"
)

// EOF is a special rune indicating the end of the buffer.
const EOF = rune(-1)

// RuneReader is a buffered reader which operates on runes.
type RuneReader struct {
	reader      *bufio.Reader
	frontBuffer *list.List
	lastRead    rune
}

// NewRuneReader creates a new instance of a RuneReader.
func NewRuneReader(reader io.Reader) *RuneReader {
	return &RuneReader{
		reader:      bufio.NewReader(reader),
		frontBuffer: list.New(),
		lastRead:    EOF,
	}
}

// Read reads and removes the next rune from the buffer and returns it.
// If there are no runes left in the buffer, it returns an EOF.
func (rr *RuneReader) Read() rune {
	var result rune
	if rr.frontBuffer.Len() > 0 {
		result = rr.frontBuffer.Remove(rr.frontBuffer.Front()).(rune)
	} else {
		var err error
		result, _, err = rr.reader.ReadRune()
		if err == io.EOF {
			result = EOF
		}
	}
	rr.lastRead = result
	return result
}

// Unread puts back the last read rune to the front of the buffer. It reverts
// the Read function. If there was no call to Read, it can't unread the rune
// and the call to this function is a no op.
func (rr *RuneReader) Unread() {
	if rr.lastRead != EOF {
		rr.frontBuffer.PushFront(rr.lastRead)
		rr.lastRead = EOF
	}
}

// Peek returns the desired amount of runes without removing them from the
// buffer. If there are not enough runes in the buffer, it returns as many
// as possible with an EOF rune at the end.
func (rr *RuneReader) Peek(n uint) ([]rune, error) {
	result := make([]rune, 0, n)
	for i := uint(0); i < n; i++ {
		read, _, err := rr.reader.ReadRune()
		if err == io.EOF {
			result = append(result, EOF)
			rr.frontBuffer.PushBack(EOF)
			return result, nil
		}
		if err != nil {
			return result, err
		}
		rr.frontBuffer.PushBack(read)
		result = append(result, read)
	}
	return result, nil
}
