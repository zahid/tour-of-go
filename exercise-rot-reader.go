package main

import (
	"io"
	"os"
	"strings"
)

type Rot13Reader struct {
	r io.Reader
}

func (r *Rot13Reader) Read(b []byte) (int, error) {
	for i := 0; i < len(b); i++ {
		if b[i] >= 'A' && b[i] <= 'M' || b[i] >= 'a' && b[i] <= 'm' {
			b[i] += 13
		}

		if b[i] >= 'N' && b[i] <= 'Z' || b[i] >= 'n' && b[i] <= 'z' {
			b[i] -= 13
		}
	}
	return len(b), nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := Rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
