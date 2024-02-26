package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)

	if err == io.EOF {
		return n, err
	}

	for i, b := range p {
		p[i] = rot13(b)
	}

	return n, err
}

func rot13(b byte) byte {
	var a byte
	switch {
	case 'a' <= b && b <= 'z':
		a = 'a'
	case 'A' <= b && b <= 'Z':
		a = 'A'
	default:
		return b
	}

	return (b-a+13)%26 + a
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
