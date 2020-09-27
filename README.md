# Peek-A-Buf

Peek-A-Buf is a Go library for a buffered reader with side effect free peeking capability.

At the moment, a buffered reader for runes is implemented meant for things like scanners.

Why is this useful if `bufio.Reader` exists? Well, the Peek function is sadly not side effect
free there:

>  Calling Peek prevents a UnreadByte or UnreadRune call from succeeding until the next read
> operation. 

But this was something needed to fix a bug in the
[SQLDumpSplitter3](https://philiplb.de/sqldumpsplitter3). So I created this little library
basically wrapping `bufio.Reader` and re-implementing small parts of it.

In addition to the side effect free peek function, it returns an `EOF` rune instead of an
error if it reaches the end of the buffer.

Have a look at the **[documentation](https://godoc.org/github.com/philiplb/peekabuf)**.

## Installation

Just `go get` it:

```bash
go get github.com/philiplb/peekabuf
```

## Usage

First, import it:

```Go
import "github.com/philiplb/peekabuf"
```

Then, create a new RuneReader with a `io.Reader` as source, in this case, simply a `bytes.BufferString` is used:

```Go
r := peekabuf.NewRuneReader(bytes.NewBufferString("pab"))
```

Now, you can read, unread, peek on it. Note the returned `EOF` rune if the end of the buffer is reached:

```Go
// read a bit in the buffer
read := r.Read() // read == 'p'
read = r.Read() // read == 'a'

// unread the last rune
r.Unread()
read = r.Read() // read == 'a'

// peek a bit in the buffer
r.Unread() // unread the last rune so we have a bit to peek
peeked, err := r.Peek(2) // peeked == {'a', 'b'}
peeked, err := r.Peek(4) // peeked == {'a', 'b', peekabuf.EOF}, note that four runes were requested

// read the rest of the buffer
read = r.Read() // read == 'a'
read = r.Read() // read == 'b'
read = r.Read() // read == peekabuf.EOF
read = r.Read() // read == peekabuf.EOF
```

## License
[MIT](https://choosealicense.com/licenses/mit/)
