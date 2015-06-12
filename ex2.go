package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy)
	for i := range a {
		a[i] = make([]uint8, dx)
	}

	for j := 0; j < dy; j++ {
		for k := 0; k < dx; k++ {
			a[j][k] = uint8(j * k)
		}
	}
	return a
}

func main() {
	pic.Show(Pic)
}
