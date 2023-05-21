package main

import (
	"fmt"

	"golang.org/x/tour/reader"
)

type EOFError []byte

func (e EOFError) Error() string {
	return fmt.Sprintf("EOF error")
}

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (int, error) {
	if b == nil || len(b) == 0 {
		return 0, EOFError(b)
	}
	for idx := range b {
		b[idx] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}

/*
Exercise: Readers

Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
*/
