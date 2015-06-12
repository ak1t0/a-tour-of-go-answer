package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannnot Sqrt negativenumber: %g", float64(e))
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return f, ErrNegativeSqrt(f)
	}

	z := float64(3)
	r := float64(0)

	for i := 0; i <= 10; i++ {
		z = z - (z*z-f)/(2*z)
		r = z
	}
	return r, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
