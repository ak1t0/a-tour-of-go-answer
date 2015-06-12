package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x := 0
	return func() int {
		x += 1
		return fiber(x - 1)
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func fiber(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	return fiber(x-1) + fiber(x-2)
}
