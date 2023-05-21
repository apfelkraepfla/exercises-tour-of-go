package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)
	for row_idx := range result {
		result[row_idx] = make([]uint8, dx)
		for col_idx := range result[row_idx] {
			result[row_idx][col_idx] = uint8(row_idx ^ col_idx)
			// result[row_idx][col_idx] = uint8((row_idx + col_idx)/2)
			// result[row_idx][col_idx] = uint8(row_idx * col_idx)
		}
	}
	return result
}

func main() {
	pic.Show(Pic)
}

/*
Exercise: Slices

Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.

(You need to use a loop to allocate each []uint8 inside the [][]uint8.)

(Use uint8(intValue) to convert between types.)

*/
