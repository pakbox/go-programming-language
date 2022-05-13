package main

import "fmt"

/* gcd calculates GCD of two integers */
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x % y
	}
	return x
}

/* fib calculates n-th Fibonacci number iteratively*/
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n ; i++ {
		x, y = y, x + y
	}
	return x
}

func main() {
	var result int
	result = gcd(4,24)
	fmt.Printf("gcd: %d\n", result)
	result = fib(10)
	fmt.Printf("fib: %d\n", result)
}
