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

## License
[MIT](https://choosealicense.com/licenses/mit/)
