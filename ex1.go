package main

import (
	"fmt"
	"math"
)

func newt(z, x float64, n int) (r float64) {
	for i := 0; i <= n; i++ {
		z = z - (z*z-x)/(2*z)
		r = z
	}
	return
}

func main() {
	fmt.Println("z = 11, x = 2,  1 loop: ", newt(11, 2, 1))
	fmt.Println("z = 11, x = 2,  5 loop: ", newt(11, 2, 5))
	fmt.Println("z = 11, x = 2, 10 loop: ", newt(11, 2, 10))
	fmt.Println("math.Sqrt             : ", math.Sqrt(2))
}
