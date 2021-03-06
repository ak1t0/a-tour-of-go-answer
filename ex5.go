package main

import (
	"fmt"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	var z complex128 = 10
	for i := 0; i < 1000000; i++ {
		z = z - (z*z*z-x)/(3*z*z)
	}
	return z
}

func main() {
	fmt.Println(Cbrt(2))
	fmt.Println(cmplx.Pow(2, 0.33333333333333333333333333333333333))
}
