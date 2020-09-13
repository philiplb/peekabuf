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
	"bytes"
	"testing"
)

func TestNewRuneReader(t *testing.T) {
	b := NewRuneReader(bytes.NewBufferString("peekabuf"))
	if b == nil {
		t.Error("expected a RuneReader pointer but got nil")
	}
}

func TestRead(t *testing.T) {
	r := NewRuneReader(bytes.NewBufferString("peekabuf"))
	expected := []rune{'p', 'e', 'e', 'k', 'a', 'b', 'u', 'f', EOF, EOF}
	for _, expectedRune := range expected {
		actualRune := r.Read()
		if actualRune != expectedRune {
			t.Errorf("expected rune to be %s but got %s", string(expectedRune), string(actualRune))
		}
	}
}

func TestUnread(t *testing.T) {
	r := NewRuneReader(bytes.NewBufferString("pab"))

	r.Unread()
	checkRune(t, r, 'p')

	r.Unread()
	checkRune(t, r, 'p')

	r.Unread()
	r.Unread()
	checkRune(t, r, 'p')

	expected := []rune{'a', 'b', EOF, EOF}
	for _, toCheck := range expected {
		r.Read()
		r.Unread()
		checkRune(t, r, toCheck)
	}
}

func TestPeek(t *testing.T) {

	r := NewRuneReader(bytes.NewBufferString("pab"))

	peeked, err := r.Peek(2)
	if err != nil {
		t.Errorf("expected no error but got %s", err)
	}
	if len(peeked) != 2 {
		t.Errorf("expected 2 peeked runes but got %d", len(peeked))
	}
	if peeked[0] != 'p' || peeked[1] != 'a' {
		t.Errorf("expected 'p' and 'a' but got '%s' and '%s'", string(peeked[0]), string(peeked[1]))
	}

	checkRune(t, r, 'p')
	checkRune(t, r, 'a')
	checkRune(t, r, 'b')
	checkRune(t, r, EOF)

	peeked, err = r.Peek(2)
	if err != nil {
		t.Errorf("expected no error but got %s", err)
	}
	if len(peeked) != 1 {
		t.Errorf("expected 1 peeked rune but got %d", len(peeked))
	}
	if peeked[0] != EOF {
		t.Errorf("expected EOL but got '%s'", string(peeked[0]))
	}

	r = NewRuneReader(bytes.NewBufferString("pab"))
	peeked, err = r.Peek(4)
	if err != nil {
		t.Errorf("expected no error but got %s", err)
	}
	if len(peeked) != 4 {
		t.Errorf("expected 4 peeked runes but got %d", len(peeked))
	}
	if peeked[0] != 'p' || peeked[1] != 'a' || peeked[2] != 'b' || peeked[3] != EOF {
		t.Errorf("expected 'p', 'a', 'b' and EOL but got '%s', '%s', '%s' and '%s'",
			string(peeked[0]), string(peeked[1]), string(peeked[2]), string(peeked[3]))
	}

	checkRune(t, r, 'p')
	checkRune(t, r, 'a')
	checkRune(t, r, 'b')
	checkRune(t, r, EOF)

}

func checkRune(t *testing.T, r *RuneReader, expected rune) {
	read := r.Read()
	if read != expected {
		t.Errorf("expected rune to be %s but got %s", string(expected), string(read))
	}
}
