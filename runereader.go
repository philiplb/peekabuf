/*
 * This file is part of the Peek-A-Buf package.
 *
 * (c) Philip Lehmann-Böhm <philip@philiplb.de>
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

package peekabuf

import (
	"bufio"
	"container/list"
	"io"
)

type RuneReader struct {
	reader      *bufio.Reader
	frontBuffer *list.List
	lastRead    rune
}

const EOF = rune(-1)

func NewRuneReader(reader io.Reader) *RuneReader {
	return &RuneReader{
		reader:      bufio.NewReader(reader),
		frontBuffer: list.New(),
		lastRead:    EOF,
	}
}

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

func (rr *RuneReader) Unread() {
	if rr.lastRead != EOF {
		rr.frontBuffer.PushFront(rr.lastRead)
		rr.lastRead = EOF
	}
}

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
