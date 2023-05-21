package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(input byte) (output byte) {

	highest_number := 0

	if 'A' <= input && input <= 'Z' {
		// ASCII numbers for capital letters: 65-90
		highest_number = 90

	} else if 'a' <= input && input <= 'z' {
		// ASCII numbers for lower letters: 97-122
		highest_number = 122
	} else {
		// anything else is not an alphabetical character
		return input
	}

	output = input + 13
	if output > byte(highest_number) {
		/* The result has to be in the range of the
		capital/lower ASCII letters. As in rot 13, we can
		only exceed the upper limit by 13, we simply can
		substract 26 to shift the outcome by the length
		of one alphabet to its respective place.
		*/
		output -= 26
	}
	return output

}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

/*
Exercise: rot13Reader

A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.

For example, the gzip.NewReader function takes an io.Reader (a stream of compressed data) and returns a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).

Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.

The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method.
*/
