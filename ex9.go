package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)
	if err != nil {
		return
	}
	for i := 0; i < n; i++ {
		p[i] = rot13er(p[i])
	}
	return
}

func rot13er(c byte) byte {
	t := c
	if c >= 'a' && c <= 'z' {
		t += 13
		if t > 'z' {
			t = t - 'z' + 'a' - 1
		}
	}
	if c >= 'A' && c <= 'Z' {
		t += 13
		if t > 'Z' {
			t = t - 'Z' + 'A' - 1
		}
	}
	return t
}

func main() {
	s := strings.NewReader(
		"Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
